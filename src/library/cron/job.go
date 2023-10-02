package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

func job() {
	c := cron.New()

	j1 := MyJob{name: "1"}
	_, errJob1 := c.AddJob("@every 5s", &j1)
	if errJob1 != nil {
		fmt.Printf("errJob1: %v", errJob1)
	}

	j2 := MyJob{name: "2"}
	_, errJob2 := c.AddJob("@every 6s", &j2)
	if errJob2 != nil {
		fmt.Printf("errJob2: %v", errJob2)
	}
	c.Start()
	select {}
}

// MyJob 对定时任务进行包装
type MyJob struct {
	name string
}

// Run 实现 Job 接口里的 Run 方法
func (j *MyJob) Run() {
	fmt.Printf("myjob : %s\n", j.name)
}
