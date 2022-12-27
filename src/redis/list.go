package main

import (
	"fmt"
	"time"
)

// list 的操作
func opList() {
	var key = "queue"
	var key1 = "queue1"
	if err := redisCli.RPush(ctx, key, "message", "message1", "message2", "message3").Err(); err != nil {
		panic(err)
	}

	// BLPop: 移出并获取列表的第一个元素， 如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止
	// redisCli.BLPop(ctx, 0, key) 使用 0 意味着永久等待，直到有这个 key
	BLPopResult, err := redisCli.BLPop(ctx, 1*time.Second, key).Result()
	if err != nil {
		panic(err)
	}

	fmt.Printf("BLPop result = %v\n", BLPopResult[0])

	// 从 source 中弹出一个元素，并插入到 dest 里，如果没元素会阻塞，所以同样有提供超时的参数
	result1, err := redisCli.BRPopLPush(ctx, key, key1, 1*time.Second).Result()
	if err != nil {
		panic(err)
	}

	fmt.Printf("BRPopLPush result = %v\n", result1)

	// 在指定 key 的某个 value 后面插入数据
	LInsertAfterResult, err := redisCli.LInsertAfter(ctx, key1, "message3", "afterMessage").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("LInsertAfter result = %v\n", LInsertAfterResult)

	// 获取长度
	lLen := redisCli.LLen(ctx, key1).Val()
	fmt.Printf("key1 len = %d\n", lLen)

	// 输出指定范围里的value
	LRangeResult, err := redisCli.LRange(ctx, key1, 0, lLen).Result()
	if err != nil {
		panic(err)
	}

	fmt.Printf("key1 value for range\n")
	for _, value := range LRangeResult {
		fmt.Printf("%s ", value)
	}
	fmt.Printf("\n")

	var (
		key2 = "queue2"
		key3 = "queue3"
	)
	// 对一个列表的头部添加元素，如果列表不存在，则会创建并添加
	LPushResult, err := redisCli.LPush(ctx, key2, "message").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("LPush result = %v\n", LPushResult)
	// 对一个已存在的列表添加元素，如果列表不存在，则不做
	LPushXResult, err := redisCli.LPushX(ctx, key3, "message").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("LPushX result = %v\n", LPushXResult)
}
