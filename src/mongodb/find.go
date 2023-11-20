package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

// 查询指定的条件
func findProjection(ctx context.Context, col *mongo.Collection) error {
	// 只返回 title，不返回 id
	option := options.Find().SetProjection(bson.D{{"title", 1}, {"_id", 0}})
	// 查询全部
	cur, err := col.Find(ctx, bson.D{}, option)
	if err != nil {
		return err
	}
	var movies = make([]Movie, 0)
	err = cur.All(ctx, &movies)
	if err != nil {
		return err
	}
	fmt.Println(movies)
	return nil
}

func findElemMatch(ctx context.Context, col *mongo.Collection) error {
	// 使用 elemMatch 对内嵌文档进行匹配，这里搜索有台湾地区的剧本
	filter := bson.D{{"areas", bson.D{{"$elemMatch", bson.D{{"locale", "zh-TW"}}}}}}
	cur, err := col.Find(ctx, filter)
	if err != nil {
		return err
	}
	var scenarios = make([]ScenarioInfo, 0)
	err = cur.All(ctx, &scenarios)
	if err != nil {
		return err
	}
	fmt.Println(scenarios)
	return nil
}

// 使用 in 条件查询
func findIn(ctx context.Context, col *mongo.Collection) error {
	// 查询 2 和 3
	cur, err := col.Find(ctx, bson.D{{"_id", bson.D{{"$in", bson.A{2, 3}}}}})
	if err != nil {
		return err
	}
	var movies = make([]Movie, 0)
	err = cur.All(ctx, &movies)
	if err != nil {
		return err
	}
	fmt.Println(movies)
	return nil
}

// 使用 nin 条件查询
func findNIn(ctx context.Context, col *mongo.Collection) error {
	// 查询不是 2 和 3
	cur, err := col.Find(ctx, bson.D{{"_id", bson.D{{"$nin", bson.A{"2", "3"}}}}})
	if err != nil {
		return err
	}
	var movies = make([]Movie, 0)
	err = cur.All(ctx, &movies)
	if err != nil {
		return err
	}
	fmt.Println(movies)
	return nil

}

func findOr(ctx context.Context, col *mongo.Collection) error {
	// 查询 2 或者 4
	cur, err := col.Find(ctx, bson.D{{"$or", bson.A{bson.D{{"_id", "2"}}, bson.D{{"_id", "4"}}}}})
	if err != nil {
		return err
	}
	var movies = make([]Movie, 0)
	err = cur.All(ctx, &movies)
	if err != nil {
		return err
	}
	fmt.Println(movies)
	return nil

}
