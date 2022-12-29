package main

import "fmt"

// HyperLogLog  是用来做基数统计的算法，HyperLogLog 的优点是，在输入元素的数量或者体积非常非常大时，计算基数所需的空间总是固定的、并且是很小的。
// 在 Redis 里面，每个 HyperLogLog 键只需要花费 12 KB 内存，就可以计算接近 2^64 个不同元素的基数。
// 因为 HyperLogLog 只会根据输入元素来计算基数，而不会储存输入元素本身，所以 HyperLogLog 不能像集合那样，返回输入的各个元素。
func opHyperLogLog() {
	var key = "hyper"
	pfAddResult, errPfAdd := redisCli.PFAdd(ctx, key, "add1", "add2", "add3").Result()
	if errPfAdd != nil {
		panic(errPfAdd)
	}
	fmt.Printf("pfAdd result = %v\n", pfAddResult)

	pfCountResult, errPfCount := redisCli.PFCount(ctx, key).Result()
	if errPfCount != nil {
		panic(errPfCount)
	}
	fmt.Printf("pfCount result = %v\n", pfCountResult)

	var key1 = "hyper1"
	pfAddResult1, errPfAdd1 := redisCli.PFAdd(ctx, key1, "1add1", "1add2", "1add3").Result()
	if errPfAdd1 != nil {
		panic(errPfAdd1)
	}
	fmt.Printf("pfAdd1 result = %v\n", pfAddResult1)

	var keyMerge = "hyperMerge"
	// 这里 merge 会将上面两部分的 els 进行合并，过滤掉重复的
	pfMergeResult, errPfMerge := redisCli.PFMerge(ctx, keyMerge, key, key1).Result()
	if errPfMerge != nil {
		panic(errPfMerge)
	}
	fmt.Printf("pfMerge result = %v\n", pfMergeResult)

	pfCountResult, errPfCount = redisCli.PFCount(ctx, keyMerge).Result()
	if errPfCount != nil {
		panic(errPfCount)
	}
	fmt.Printf("merge pfCount result = %v\n", pfCountResult)
}
