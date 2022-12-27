package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func init() {
	redisCli = redis.NewClient(&redis.Options{
		Addr:        "localhost:6379",
		Password:    "", // no password set
		DB:          0,  // use default DB
		ReadTimeout: -1,
	})

	pong, err := redisCli.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Redis pong = ", pong)
}

func main() {

	//basicUsage()
	//customCmd()
	//opList()
	//opString()
	//opPipelined()
	opTxPipelined()
}
