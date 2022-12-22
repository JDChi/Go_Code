package main

import (
	"fmt"
	"regexp"
)

func main() {
	var plainText = "hello, my name is Ben, Please visit my website at https://www.forta.com/."

	// 普通文本的匹配
	// 由于是普通文本的匹配，所以会区分大小写
	plainTextRegexp := regexp.MustCompile("Ben")
	if plainTextRegexp == nil {
		panic("plainTextRegexp err")
	}
	fmt.Println(".......single match.............")
	// 这里只会返回一个
	plainTextResult := plainTextRegexp.FindStringSubmatch(plainText)
	fmt.Println(plainTextResult)
	plainTextRegexp = regexp.MustCompile("my")
	if plainTextRegexp == nil {
		panic("plainTextRegexp err")
	}
	fmt.Println(".......multiple match.............")
	// 这里会匹配到多个来返回
	plainTextResult1 := plainTextRegexp.FindAllStringSubmatch(plainText, -1)
	fmt.Println(plainTextResult1)

}
