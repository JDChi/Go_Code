package main

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func insertOne(ctx context.Context, col *mongo.Collection) error {
	movie := Movie{
		ID:    "2",
		Title: "Stand by me",
	}
	_, errInsert := col.InsertOne(ctx, movie)
	return errInsert
}

func insertMany(ctx context.Context, col *mongo.Collection) error {
	movie1 := Movie{ID: "3", Title: "E.T."}
	movie2 := Movie{ID: "4", Title: "Blade Runner"}

	_, err := col.InsertMany(ctx, []any{
		movie1,
		movie2,
	})
	return err
}
