package main

import (
	"github.com/gofiber/fiber"
)

var defaultCert keyPairPath = keyPairPath{
	Cert: "/etc/letsencrypt/live/rincons.cc/fullchain.pem",
	Key:  "/etc/letsencrypt/live/rincons.cc/privkey.pem",
}

var filesPath string = "home/jeff/Documents/go/src/github.com/rincons/files"

var tpl404Path string = filesPath + "/tpl/404.html"

var appConfig fiber.Config = fiber.Config{
	ServerHeader: "rincons.cc",
}

//DB configs
const (
	dbUser string = "jeff"
	dbPwd  string = "1040"
	dbDb   string = "rebbachi"
)
