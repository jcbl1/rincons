package main

import (
	"context"
	// "fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
)

type AutoReplyPost struct {
	MsgIn, MsgOut string
}

func Conn(msgin string) string {
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

	collection := client.Database("rebbachi").Collection("autoReply")

	filter := bson.D{{"MsgIn", msgin}}
	var post AutoReplyPost

	err = collection.FindOne(context.TODO(), filter).Decode(&post)

	// insertResult, err := collection.InsertOne(context.TODO(), Post{Title: "Haha", Body: "heihei"})
	// if err != nil {
	// 	panic(err.Error())
	// }

	// fmt.Println("Inserted post with ID:", insertResult.InsertedID)

	// fmt.Println("Found post with MsgIn ", post.MsgIn)

	return post.MsgOut
}
