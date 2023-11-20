package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoCli *mongo.Client

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var err error
	mongoCli, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://admin:123456@localhost:27017"))
	if err != nil {
		panic(err)
	}

}

func main() {
	ctx := context.Background()
	defer func() {
		if err := mongoCli.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	// 获取数据库
	testDB := mongoCli.Database("test")

	person1Col := testDB.Collection("person1")

	// err := insertOne(ctx, person1Col)
	// err := insertMany(ctx, person1Col)
	// err := findOne(ctx, "3", person1Col)
	// err := find(ctx, person1Col)
	err := updateOne(ctx, person1Col, "2")
	if err != nil {
		fmt.Println(err)
		return
	}

}

type Movie struct {
	ID    string `bson:"_id"`
	Title string `bson:"title"`
}
