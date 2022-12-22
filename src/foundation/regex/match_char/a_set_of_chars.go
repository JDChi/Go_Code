package main

import (
	"fmt"
	"regexp"
)

func main() {
	// 一组字符的匹配
	plainText := `sales.xls
		             sales1.xls
                     orders3.xls
                     sales1.xls
                     sales2.xls
                     sam.xls
                     na1.xls
                     na2.xls
                     sa1.xls
                     ca1.xls`
	// 用[]来指定集合必须匹配的某些条件
	plainTextRegexp := regexp.MustCompile("[ns]a.\\.xls")
	if plainTextRegexp == nil {
		panic("plainTextRegexp err")
	}
	fmt.Println(".......a set of match.............")
	plainTextResult1 := plainTextRegexp.FindAllStringSubmatch(plainText, -1)
	fmt.Println(plainTextResult1)

	// 利用字符集合区间，只过滤出带有数字的文件
	plainTextRegexp = regexp.MustCompile("[ns]a[0-9]\\.xls")
	if plainTextRegexp == nil {
		panic("plainTextRegexp err")
	}
	fmt.Println("....... a set of limit match.............")
	plainTextResult1 = plainTextRegexp.FindAllStringSubmatch(plainText, -1)
	fmt.Println(plainTextResult1)

	// 使用 ^ 来进行排除
	plainTextRegexp = regexp.MustCompile("[ns]a[^0-9]\\.xls")
	if plainTextRegexp == nil {
		panic("plainTextRegexp err")
	}
	fmt.Println(".......exclude a set of match.............")
	plainTextResult1 = plainTextRegexp.FindAllStringSubmatch(plainText, -1)
	fmt.Println(plainTextResult1)

}
