package main

import (
	"crypto/sha1"
	"fmt"
	"github.com/gofiber/fiber"
	// "os"
	"sort"
	// "strconv"
	"strings"
	// "log"
	// "go.mongodb.org/mongo-driver/bson"
)

type Re struct {
	ToUserName   string
	FromUserName string
	CreateTime   int
	MsgType      string
	Content      string
	MsgId        int64
}

func rebbachiCom(c *fiber.Ctx) error {
	// f, err := os.OpenFile("/home/duono/tmpXX", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer f.Close()

	// echostr := c.FormValue("echostr")
	if c.Method() == "POST" {

		tplStr := "<xml><ToUserName>%s</ToUserName><FromUserName>%s</FromUserName><CreateTime>%d</CreateTime><MsgType>text</MsgType><Content>%s</Content></xml>"
		re := new(Re)
		if err := c.BodyParser(re); err != nil {
			return err
		}

		msgout := Conn(re.Content)
		if msgout != "" {
			return c.SendString(fmt.Sprintf(tplStr, re.FromUserName, re.ToUserName, re.CreateTime, msgout))
		}

		// f.WriteString(re.ToUserName + "\n" + re.FromUserName + "\n" + re.Content + "\n" + strconv.Itoa(re.CreateTime) + "\n")
		return c.SendString(fmt.Sprintf(tplStr, re.FromUserName, re.ToUserName, re.CreateTime, re.Content+"是个好人"))
	}

	// 以下是验证Token
	echostr := c.FormValue("echostr")
	if echostr != "" {
		signature := c.FormValue("signature")
		timestamp := c.FormValue("timestamp")
		nonce := c.FormValue("nonce")
		token := "sdkfh2lsdkfAl"

		// f, err := os.Create("/home/duono/tmpPP")
		// if err != nil {
		// 	panic(err.Error())
		// }
		// defer f.Close()

		tmpSlice := []string{timestamp, nonce, token}
		sort.Strings(tmpSlice)

		str := strings.Join(tmpSlice, "")
		sign := fmt.Sprintf("%x", sha1.Sum([]byte(str)))

		if sign == signature {
			return c.SendString(echostr)
		}
		return c.SendString("not match")

		// echostr := c.Get("echostr")
		// return c.SendString(echostr)
	}
	return c.SendString("网站正在建设中...")
}
