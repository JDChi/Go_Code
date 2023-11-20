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

func insertMany1(ctx context.Context, col *mongo.Collection) error {
	scenario := ScenarioInfo{Id: 3, Areas: []Area{
		{Locale: "en", Title: "en_title", Introduction: "en_intro"},
		{Locale: "zh", Title: "中文标题", Introduction: "中文介绍"},
	}}
	scenario1 := ScenarioInfo{Id: 4, Areas: []Area{
		{Locale: "en", Title: "en_title1", Introduction: "en_intro_1"},
		{Locale: "zh-TW", Title: "中文标题1", Introduction: "中文介绍1"},
	}}
	scenario2 := ScenarioInfo{Id: 5, Areas: []Area{
		{Locale: "en", Title: "en_title2", Introduction: "en_intro_2"},
		{Locale: "jp", Title: "日文标题2", Introduction: "日文介绍2"},
	}}

	_, err := col.InsertMany(ctx, []any{
		scenario,
		scenario1,
		scenario2,
	})
	return err
}
