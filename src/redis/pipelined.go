package main

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

func opPipelined() {
	// 普通执行 100 次命令和使用管道来执行，管道的效率会高些
	var start = time.Now()
	// 普通执行 100 次命令
	for i := 0; i < 100; i++ {
		redisCli.Ping(ctx)
	}
	var end = time.Since(start)
	fmt.Printf("normal end = %v\n", end)

	start = time.Now()
	// 使用管道来执行 100 次命令
	_, err := redisCli.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		for i := 0; i < 100; i++ {
			redisCli.Ping(ctx)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	end = time.Since(start)
	fmt.Printf("pipelined end = %v\n", end)
}
