package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func basicUsage() {
	var ctx = context.Background()
	var key = "go_code_basic"

	getResult, errGet := redisCli.Get(ctx, "go_code_basic1").Result()
	if errors.Is(errGet, redis.Nil) {
		fmt.Printf("key is not exist\n")
	}

	setResult, errSet := redisCli.Set(ctx, key, "hello", -1).Result()
	if errSet != nil {
		panic(errSet)
	}
	fmt.Printf("setResult: %v\n", setResult)

	getResult, errGet = redisCli.Get(ctx, key).Result()
	if errGet != nil {
		panic(errGet)
	}
	fmt.Printf("getResult: %v\n", getResult)
}
