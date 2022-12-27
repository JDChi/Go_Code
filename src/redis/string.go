package main

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

func opString() {
	var key = "page"

	setResult, err := redisCli.Set(ctx, key, 1, redis.KeepTTL).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("set result = %v\n", setResult)

	// 给 key 存储的数字值加 1，如果这个key的值不是数字，就会得到错误
	incrResult, err := redisCli.Incr(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Incr result = %v\n", incrResult)

	getResult, err := redisCli.Get(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("get result = %v\n", getResult)

	// 给 key 存储的数字值加指定的数值
	incrByResult, err := redisCli.IncrBy(ctx, key, 10).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("IncrBy result = %v\n", incrByResult)

	getResult, err = redisCli.Get(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("get result = %v\n", getResult)

	var nxKey = "nxKey"
	// key 不存在的时候设置值，成功返回 true，失败返回 false，意味着 key 已经存在，通常用来做加锁，解锁即把 key 删掉
	nxResult, err := redisCli.SetNX(ctx, nxKey, "nx", 1*time.Hour).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("nx result = %v\n", nxResult)

}
