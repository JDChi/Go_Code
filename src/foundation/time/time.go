package main

import (
	"fmt"
	"time"
)

func main() {
	/////////////////////// timezone
	nowTime := time.Now()
	fmt.Printf("nowTime = %v\n", nowTime)
	// 印尼时区
	jakartaTimezone := time.FixedZone("Asia/Jakarta", 7*3600)
	// 将当前的 unix 转成印尼时间，即现在印尼那边的时间
	jakartaTime := nowTime.In(jakartaTimezone)
	// 印尼零点的处理
	jakartaZeroTime := time.Date(jakartaTime.Year(), jakartaTime.Month(), jakartaTime.Day(),
		0, 0, 0, 0,
		jakartaTimezone)
	fmt.Printf("jakartaTime = %v , jakartaZeroTime = %v\n", jakartaTime, jakartaZeroTime)
	////////////////// add
	nextTime := nowTime.Add(1 * time.Hour)
	// 如果是在外面定义的数字，就需要转成 time.Duration 来处理
	var hour = 2
	nextTime2 := nowTime.Add(time.Duration(hour) * time.Hour)
	fmt.Printf("nextTime = %v, nextTime2 = %v", nextTime, nextTime2)
}
