package main

import (
	"crypto/tls"
	"github.com/gofiber/fiber"
	"net"
	"os"

	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type keyPairPath struct {
	Cert, Key string
}

func main() {
	app := fiber.New(appConfig)

	ln := newListener()

	routing(app)

	go app.Listen(":80")
	app.Listener(ln)
}

func newListener() net.Listener {
	ln, err := net.Listen("tcp", ":443")
	if err != nil {
		panic(err.Error())
	}
	cer, err := tls.LoadX509KeyPair(defaultCert.Cert, defaultCert.Key)
	if err != nil {
		panic(err.Error())
	}
	return tls.NewListener(ln, &tls.Config{
		Certificates: []tls.Certificate{cer},
	})
}

type Visit struct {
	IP, LastTime string
}

func rinconsCc(c *fiber.Ctx) error {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://" + dbUser + ":" + dbPwd + "@localhost:27017/" + dbDb))
	if err != nil {
		panic(err.Error())
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err.Error())
	}
	defer client.Disconnect(ctx)
	collection := client.Database("rebbachi").Collection("visits")

	filter := bson.D{{"ip", c.IP()}}
	var v Visit
	err = collection.FindOne(context.TODO(), filter).Decode(&v)
	if v == (Visit{}) {
		_, err = collection.InsertOne(context.TODO(), Visit{c.IP(), time.Now().Format("Jan 2 15:04:05, 2006")})
		if err != nil {
			panic(err.Error())
		}
	}

	return c.SendFile(filesPath + "/tpl/homepage.html")
}

func serveFile(c *fiber.Ctx, d string) error {
	if c.Protocol() == "http" {
		c.Redirect("https://" + c.Hostname() + c.OriginalURL())
		return c.SendStatus(fiber.StatusSeeOther)
	} else {
		switch c.Hostname() {
		case "rincons.cc":
			_, err := os.Stat(filesPath + d + c.Params("filename"))
			if !os.IsNotExist(err) {
				return c.SendFile(filesPath + d + c.Params("filename"))
			} else {
				c.SendFile(tpl404Path)
				return c.SendStatus(fiber.StatusNotFound)
			}
		case "www.rincons.cc":
			c.Redirect("https://rincons.cc" + c.OriginalURL())
			return c.SendStatus(fiber.StatusSeeOther)
		default:
			c.SendFile(tpl404Path)
			return c.SendStatus(fiber.StatusNotFound)
		}
	}
	return fiber.NewError(996, "Unexpected error.")
}
