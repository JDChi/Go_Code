package main

import (
	"fmt"
	"regexp"
)

func main() {
	// 多个字符的匹配
	plainText := `sales.xls
		             sales1.xls
                     orders3.xls
                     sales1.xls
                     sales2.xls 
                     na1.xls
                     na2.xls
                     sa1.xls`
	// . 代表着匹配任意字符，包括 . 本身
	plainTextRegexp := regexp.MustCompile("sales.")
	if plainTextRegexp == nil {
		panic("plainTextRegexp err")
	}
	fmt.Println(".......any char match.............")
	plainTextResult1 := plainTextRegexp.FindAllStringSubmatch(plainText, -1)
	fmt.Println(plainTextResult1)
}
