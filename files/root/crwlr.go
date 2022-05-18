package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gocolly/colly"
)

func main() {
	// 打开或者创建文件，文件名为运行本程序时伴随输入的参数
	f, err := os.OpenFile(os.Args[1], os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()

	dura, err := time.ParseDuration("300ms")
	if err != nil {
		panic(err.Error())
	}

	// 运用Colly包定义两个收集器（collector），collector负责收集页面1（https://gongyi.weibo.com/list/personal)的各个项目的链接，collector1负责在项目页面收集需要的信息
	collector := colly.NewCollector(colly.AllowedDomains("gongyi.weibo.com"))
	collector1 := colly.NewCollector(colly.AllowedDomains("gongyi.weibo.com"))

	collector1.OnRequest(func(r *colly.Request) {
		fmt.Println(r.URL)
	})

	// 在（每个）DOM匹配class="tit"时，collector1执行以下函数，counter1用来避免多次匹配
	counter1 := true
	collector1.OnHTML(".tit", func(e *colly.HTMLElement) {
		if counter1 {
			//写入文件
			f.WriteString("项目标题：" + e.ChildText("strong") + " ")
			counter1 = false
		}
	})

	// 在（每个）DOM匹配DOM名为dd时，collector1执行以下函数，counter用来指示收集的信息的类别，以及避免重复收集
	counter := 0
	collector1.OnHTML("dd", func(e *colly.HTMLElement) {
		i := e.ChildText("i")
		if i != "" {
			// 第一次匹配是写入“项目筹款额”，第二次“捐款人数”，第三次“目标筹款额”
			switch counter {
			case 0:
				f.WriteString("项目筹款额：" + e.ChildText("i") + " ")
				counter++
			case 1:
				f.WriteString("捐款人数：" + e.ChildText("i") + " ")
				counter++
			case 2:
				f.WriteString("目标筹款额：" + e.ChildText("i") + " ")
				counter++
			}
		}
	})

	// 在（每个）DOM匹配class="title"时，collector执行以下函数
	collector.OnHTML(".title", func(element *colly.HTMLElement) {
		href := element.ChildAttr("a", "href")
		// 如果该DOM的href不为空的话，则让collector1访问相应的网址
		if href != "" {
			time.Sleep(dura)
			collector1.Visit("https://gongyi.weibo.com" + element.ChildAttr("a", "href"))
			f.WriteString("\n")
			// 初始化counter和counter1
			counter = 0
			counter1 = true
		}
	})
	// 根据运行本程序时伴随输入的参数，决定要爬取多少页的数据
	if len(os.Args) == 2 {
		for page := 1; page <= 20; page++ {
			collector.Visit("https://gongyi.weibo.com/list/personal?page=" + strconv.Itoa(page))
			time.Sleep(dura)
		}
	} else {
		pageLim, err := strconv.Atoi(os.Args[2])
		if err != nil {
			panic(err.Error())
		}
		for page := 1; page <= pageLim; page++ {
			collector.Visit("https://gongyi.weibo.com/list/personal?page=" + strconv.Itoa(page))
			time.Sleep(dura)
		}
	}
}
