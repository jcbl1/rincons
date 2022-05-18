package main

import (
	"fmt"
	"github.com/gofiber/fiber"
	"os"
	"strings"
)

func routing(app *fiber.App) {
	// 一般对根目录的请求
	app.Get("/", func(c *fiber.Ctx) error {
		if c.Protocol() == "http" {
			c.Redirect("https://" + c.Hostname())
			return c.SendStatus(fiber.StatusSeeOther)
		} else {
			switch c.Hostname() {
			case "rincons.cc":
				err := rinconsCc(c)
				return err
			case "rebbachi.com":
				// err := rebbachiCom(c, []string{
				// 	c.Get("signature"),
				// 	c.Get("timestamp"),
				// 	c.Get("nonce"),
				// 	c.Get("echostr"),
				// })
				err := rebbachiCom(c)
				return err
			default:
				if strings.HasPrefix(c.Hostname(), "www.") {
					c.Redirect("https://" + strings.TrimPrefix(c.Hostname(), "www."))
					return c.SendStatus(fiber.StatusSeeOther)
				}
				c.SendFile(tpl404Path)
				return c.SendStatus(fiber.StatusNotFound)
			}
		}
		return fiber.NewError(996, "Unexpected error.")
	})
	//rebbachi.com 微信自动回复
	app.Post("/", func(c *fiber.Ctx) error {
		if c.Protocol() == "http" {
			c.Redirect("https://" + c.Hostname())
			return c.SendStatus(fiber.StatusSeeOther)
		}
		switch c.Hostname() {
		case "rebbachi.com":
			err := rebbachiCom(c)
			return err
		case "www.rebbachi.com":
			c.Redirect("https://rebbachi.com")
			return c.SendStatus(fiber.StatusSeeOther)
		default:
			c.SendFile(tpl404Path)
			return c.SendStatus(fiber.StatusNotFound)
		}
		return fiber.NewError(996, "Unexpected error.")

	})

	app.Get("/v/:filename", func(c *fiber.Ctx) error {
		return serveFile(c, "/videos/")
	})
	app.Get("/jquery/:filename", func(c *fiber.Ctx) error {
		return serveFile(c, "/jquery/")
	})
	app.Get("/p/:filename", func(c *fiber.Ctx) error {
		return serveFile(c, "/pictures/")
	})
	app.Get("/css/:filename", func(c *fiber.Ctx) error {
		return serveFile(c, "/tpl/css/")
	})
	app.Get("/r/:filename", func(c *fiber.Ctx) error {
		return serveFile(c, "/root/")
	})
	//文件上传
	app.Get("/uploader", func(c *fiber.Ctx) error {
		if c.Protocol() == "http" {
			c.Redirect("https://" + c.Hostname() + c.OriginalURL())
			return c.SendStatus(fiber.StatusSeeOther)
		}
		switch c.Hostname() {
		case "rincons.cc":
			return c.SendFile(filesPath + "/tpl/uploader.html")
		case "www.rincons.cc":
			c.Redirect("https://rincons.cc/uploader")
			return c.SendStatus(fiber.StatusSeeOther)
		default:
			c.SendFile(tpl404Path)
			return c.SendStatus(fiber.StatusNotFound)
		}
		return fiber.NewError(996, "Unexpected error.")
	})
	app.Post("/uploader", func(c *fiber.Ctx) error {
		form, err := c.MultipartForm()
		if err != nil {
			return err
		}
		files := form.File["documents"]
		rtrStr := ""
		for _, file := range files {
			_, err = os.Stat(filesPath + fmt.Sprintf("/root/%v", file.Filename))
			if !os.IsNotExist(err) {
				rtrStr += fmt.Sprintf("%v", file.Filename) + " "
			} else {
				c.SaveFile(file, fmt.Sprintf(filesPath+"/root/%v", file.Filename))
			}
		}
		if rtrStr != "" {
			return c.SendString("上传成功\n" + rtrStr + "已存在，无法上传！")
		}
		return c.SendString("上传成功！")
	})

	app.Get("/zohoverify/:filename", func(c *fiber.Ctx) error {
		if c.Protocol() == "http" {
			c.Redirect("https://" + c.Hostname() + c.OriginalURL())
			return c.SendStatus(fiber.StatusSeeOther)
		}
		d := "/tmp/"
		switch c.Hostname() {
		case "rebbachi.com":
			_, err := os.Stat(filesPath + d + c.Params("filename"))
			if !os.IsNotExist(err) {
				return c.SendFile(filesPath + d + c.Params("filename"))
			} else {
				c.SendFile(tpl404Path)
				return c.SendStatus(fiber.StatusNotFound)
			}
		case "www.rebbachi.com":
			c.Redirect("https://rebbachi.com" + c.OriginalURL())
			return c.SendStatus(fiber.StatusSeeOther)
		default:
			c.SendFile(tpl404Path)
			return c.SendStatus(fiber.StatusNotFound)
		}
		return fiber.NewError(996, "Unexpected error.")
	})

	//404
	app.Use(func(c *fiber.Ctx) error {
		if c.Protocol() == "http" {
			c.Redirect("https://" + c.Hostname() + c.OriginalURL())
			return c.SendStatus(fiber.StatusSeeOther)
		}
		if strings.HasPrefix(c.Hostname(), "www.") {
			c.Redirect("https://" + strings.TrimPrefix(c.Hostname(), "www.") + c.OriginalURL())
			return c.SendStatus(fiber.StatusSeeOther)
		}
		c.SendFile(tpl404Path)
		return c.SendStatus(fiber.StatusNotFound)
	})

}
