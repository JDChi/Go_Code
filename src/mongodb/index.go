package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func createIndex(ctx context.Context, col *mongo.Collection) error {
	indexModel := mongo.IndexModel{Keys: bson.D{{"name", 1}}}
	name, err := col.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		return err
	}
	fmt.Println(name)
	return nil
}

func dropIndex(ctx context.Context, col *mongo.Collection) error {
	raw, err := col.Indexes().DropOne(ctx, "name_1")
	if err != nil {
		return err
	}
	fmt.Println(raw)
	return nil
}
