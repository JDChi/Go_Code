package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func findOne(ctx context.Context, id string, col *mongo.Collection) error {
	var movie Movie
	filter := bson.D{{"_id", id}}
	err := col.FindOne(ctx, filter).Decode(&movie)
	if err != nil {
		return err
	}
	fmt.Println(movie)
	return nil
}

func find(ctx context.Context, col *mongo.Collection) error {
	// 想要找全部的，也要给一个空的bson.D
	cur, err := col.Find(ctx, bson.D{})
	if err != nil {
		return err
	}

	var movies = make([]Movie, 0)
	err = cur.All(ctx, &movies)
	if err != nil {
		return err
	}
	fmt.Println(movies)

	for cur.Next(ctx) {
		var movie Movie
		err := cur.Decode(&movie)
		if err != nil {
			return err
		}
		movies = append(movies, movie)
	}
	fmt.Println(movies)
	return nil
}
