package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func timezone() {
	c := cron.New()

	// 在东京时区，每天的 4 点 30 分运行
	_, err := c.AddFunc("CRON_TZ=Asia/Tokyo 30 04 * * *", func() {
		fmt.Println("CRON_TZ=Asia/Tokyo 30 04 * * *")
	})
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	c.Start()
	defer c.Stop()

	// 在 cron 初始化的时候指定时区
	tc, _ := time.LoadLocation("Asia/Tokyo")
	c1 := cron.New(cron.WithLocation(tc))
	_, err1 := c1.AddFunc("30 04 * * *", func() {
		fmt.Println("30 04 * * *")
	})
	if err1 != nil {
		fmt.Printf("err1: %v\n", err1)
	}

	c1.Start()
	defer c1.Stop()

	select {}
}
