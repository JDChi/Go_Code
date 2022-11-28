package main

import (
	"fmt"
	"time"
)

func main() {
	// 时区转换
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
	fmt.Printf("jakartaTime = %v , jakartaZeroTime = %v", jakartaTime, jakartaZeroTime)

}
