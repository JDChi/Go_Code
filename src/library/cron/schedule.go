package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func schedule() {
	c := cron.New()

	c.Schedule(&MySchedule{}, cron.FuncJob(func() {
		fmt.Println("hello")
	}))

	c.Start()

	select {}

}

type MySchedule struct {
}

// Next 实现 Schedule 接口的 Next 方法
// Start 的时候，会触发 Next 方法，我们可以在这个的基础上，返回下一次的运行时间
// 之后每次触发都会调用 Next，我们就可以在这里做文章，动态修改下一次的触发时间
func (s *MySchedule) Next(t time.Time) time.Time {
	fmt.Printf("now time: %v\n", t)
	result := t.Add(5 * time.Second)
	return result
}
