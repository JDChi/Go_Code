package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

func linux() {
	c := cron.New()

	// 每1分钟执行，等同于 */1 * * * * 和 @every 1m
	_, err := c.AddFunc("* * * * *", func() {
		fmt.Println("* * * * *")
	})
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	// 每 5 分钟执行，等同于 @every 5m
	// / 表示频率
	_, err1 := c.AddFunc("*/5 * * * *", func() {
		fmt.Println("* * * * *")
	})
	if err1 != nil {
		fmt.Printf("err1: %v\n", err1)
	}

	// 每个小时的 01 分执行
	_, err2 := c.AddFunc("1 * * * *", func() {
		fmt.Println("1 * * * *")
	})
	if err2 != nil {
		fmt.Printf("err2: %v\n", err2)
	}

	// 每个月的第一天的 0 点 0 分运行，相当于 @monthly
	_, err3 := c.AddFunc("0 0 1 * *", func() {
		fmt.Println("0 0 1 * *")
	})
	if err3 != nil {
		fmt.Printf("err3: %v\n", err3)
	}

	// 每周的第一天（周六或周日）的 0 点 0 分运行，相当于 @weekly
	// 像一些发奖需求，在每周一进行发奖，就可以写成 0 0 * * 1
	_, err4 := c.AddFunc("0 0 * * 0", func() {
		fmt.Println("0 0 * * 0")
	})
	if err4 != nil {
		fmt.Printf("err4: %v\n", err4)
	}

	// 每天的0 点 0 分运行一次，相当于 @daily 和 @midnight
	_, err5 := c.AddFunc("0 0 * * *", func() {
		fmt.Println("0 0 * * *")
	})
	if err5 != nil {
		fmt.Printf("err5: %v\n", err5)
	}

	// 每年运行一次，在 1 月 1 日 0 点 0 分，相当于 @yearly @annually
	_, err6 := c.AddFunc("0 0 1 1 *", func() {
		fmt.Println("0 0 1 1 *")
	})
	if err6 != nil {
		fmt.Printf("err6: %v\n", err6)
	}

	// 每小时运行一次，相当于 @hourly
	_, err7 := c.AddFunc("0 * * * *", func() {
		fmt.Println("0 * * * *")
	})
	if err7 != nil {
		fmt.Printf("err7: %v\n", err7)
	}

	// - 表示范围 , 表示列表
	// 这里的意思是，在每天的 3 点到 6 点和 20 点到 23 点这两个范围内，每个 30 分执行一次
	// 也就是 3 点 30，4 点 30，5 点 30 这样
	_, err8 := c.AddFunc("30 3-6,20-23 * * *", func() {
		fmt.Println("30 3-6,20-23 * * *")
	})
	if err8 != nil {
		fmt.Printf("err8: %v\n", err8)
	}

	c.Start()
	defer c.Stop()

	select {}
}
