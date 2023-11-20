package main

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func updateOne(ctx context.Context, col *mongo.Collection, id string) error {
	col.Indexes()

	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"title", "The Matrix"}}}}
	_, err := col.UpdateOne(ctx, filter, update)
	return err
}
