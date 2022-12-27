package main

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func opTxPipelined() {
	// 事务操作，在 redis 里，是使用 Multi 开头，编写多条命令，Exec 结尾
	// 在 go-redis 里，是使用 txPipelined，类似与 Pipelined，但在前后包装了 Multi 和 Exec
	var txPipelinedKey = "txPipelinedKey"
	_, errTxPipelined := redisCli.TxPipelined(ctx, func(pipeliner redis.Pipeliner) error {
		pipeliner.Incr(ctx, txPipelinedKey)
		pipeliner.Incr(ctx, txPipelinedKey)
		pipeliner.Incr(ctx, txPipelinedKey)
		return nil
	})

	if errTxPipelined != nil {
		// 事务执行失败有特定的错误
		if errors.Is(errTxPipelined, redis.TxFailedErr) {
			fmt.Printf("txFailedErr\n")
			return
		}

		panic(errTxPipelined)
	}

	getResult, errGet := redisCli.Get(ctx, txPipelinedKey).Result()
	if errGet != nil {
		panic(errGet)
	}
	fmt.Printf("get result = %v\n", getResult)

	fmt.Printf("watch\n")
	var watchKey = "watchKey"
	// 事务搭配 Watch，可以监视 key，如果在事务执行之前，key 的值被修改了，则事务执行失败
	// 具体的应用场景在一些先读后写的情况下，如需要查询某个值，然后再进行修改
	errWatch := redisCli.Watch(ctx, func(tx *redis.Tx) error {
		getResult, errGet := redisCli.Get(ctx, watchKey).Result()
		if errGet != nil && !errors.Is(errGet, redis.Nil) {
			return errGet
		}
		fmt.Printf("get result = %v\n", getResult)

		// 如果在这里对key进行了修改，则会导致事务执行失败
		//_, errSet := redisCli.Set(ctx, watchKey, 10, time.Second).Result()
		//if errSet != nil {
		//	return errSet
		//}

		_, errTx := tx.TxPipelined(ctx, func(pipeliner redis.Pipeliner) error {
			// 做具体的事务操作
			pipeliner.Incr(ctx, watchKey)
			pipeliner.Incr(ctx, watchKey)
			return nil
		})
		if errTx != nil {
			return errTx
		}

		return nil
	}, watchKey) // 记得这里要加上 watch 的 key
	if errWatch != nil {
		// 判断到是事务操作失败的操作，做具体的处理
		if errors.Is(errWatch, redis.TxFailedErr) {
			fmt.Printf("txFailedErr\n")
			return
		}
		panic(errWatch)
	}

	getResult, errGet = redisCli.Get(ctx, watchKey).Result()
	if errGet != nil {
		panic(errGet)
	}
	fmt.Printf("watch get result = %v\n", getResult)

}
