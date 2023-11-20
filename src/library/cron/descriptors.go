package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

func descriptors() {
	c := cron.New()
	// @every
	// 支持秒级别，会在启动 6s 后执行
	_, errEveryS := c.AddFunc("@every 6s", func() {
		fmt.Println("every 6s")
	})
	if errEveryS != nil {
		fmt.Printf("errEveryS: %v\n", errEveryS)
	}
	// 0.1 min = 6s
	_, errEveryM := c.AddFunc("@every 0.1m", func() {
		fmt.Println("every 0.1m")
	})

	if errEveryM != nil {
		fmt.Printf("errEveryM: %v\n", errEveryM)
	}

	// @yearly @annually 每年执行一次，即 1 月 1 日的时候开始执行
	_, errYearly := c.AddFunc("@yearly", func() {
		fmt.Println("yearly")
	})
	if errYearly != nil {
		fmt.Printf("errYearly: %v\n", errYearly)
	}

	_, errAnnually := c.AddFunc("@annually", func() {
		fmt.Println("annually")
	})
	if errAnnually != nil {
		fmt.Printf("errAnnually: %v\n", errAnnually)
	}

	// @monthly 每个月的第一天的零点执行
	_, errMonthly := c.AddFunc("@monthly", func() {
		fmt.Println("monthly")
	})
	if errMonthly != nil {
		fmt.Printf("errMonthly: %v\n", errMonthly)
	}

	// @weekly 每周的第一天零点执行，这个第一天可能是周六，也可能是周日
	_, errWeekly := c.AddFunc("@weekly", func() {
		fmt.Println("weekly")
	})
	if errWeekly != nil {
		fmt.Printf("errWeekly: %v\n", errWeekly)
	}

	// @hourly 每小时执行一次，启动后一小时执行
	_, errHourly := c.AddFunc("@hourly", func() {
		fmt.Println("hourly")
	})
	if errHourly != nil {
		fmt.Printf("errHourly: %v\n", errHourly)
	}

	// @daily @midnight，每天执行一次，在每天的零点执行
	_, errDaily := c.AddFunc("@daily", func() {
		fmt.Println("daily")
	})
	if errDaily != nil {
		fmt.Printf("errDaily: %v\n", errDaily)
	}

	_, errMidnight := c.AddFunc("@midnight", func() {
		fmt.Println("midnight")
	})
	if errMidnight != nil {
		fmt.Printf("errMidnight: %v\n", errMidnight)
	}

	c.Start()
	defer c.Stop()

	select {}
}
