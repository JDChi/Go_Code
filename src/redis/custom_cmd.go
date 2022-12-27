package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

// 自定义命令
func customCmd() {
	var key = "go_code_basic"
	ctx := context.Background()

	Get := func(ctx context.Context, key string) *redis.StringCmd {
		cmd := redis.NewStringCmd(ctx, "get", key)
		redisCli.Process(ctx, cmd)
		return cmd
	}

	v, err := Get(ctx, key).Result()
	if err != nil {
		panic(err)
	}

	fmt.Printf("custom cmd result = %v\n", v)

	result, err := redisCli.Do(ctx, "get", key).Result()
	if err != nil {
		panic(err)
	}

	fmt.Printf("do result = %v\n", result)
}
