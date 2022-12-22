package main

import (
	"fmt"
	"regexp"
)

func main() {
	// match num or non num
	var plainText = `var myArray = new Array();
                     ...
                     if (myArray[0] == (0) {
                     ...
                     }`
	// \d 表示匹配数字 \D 表示匹配非数字
	plainTextRegexp := regexp.MustCompile("myArray\\[\\d\\]")
	if plainTextRegexp == nil {
		panic("plainTextRegexp err")
	}
	fmt.Println(".......match digit or not.............")
	plainTextResult := plainTextRegexp.FindAllStringSubmatch(plainText, -1)
	fmt.Println(plainTextResult)

	plainText = `11213
                 A1C2E3
                 48075
                 48237
                 M1B4F2
                 90046
                 H1H2H2`
	// match char
	plainTextRegexp = regexp.MustCompile("\\w\\d\\w\\d\\w\\d")
	if plainTextRegexp == nil {
		panic("plainTextRegexp err")
	}
	fmt.Println(".......match letter/digit or not.............")
	plainTextResult = plainTextRegexp.FindAllStringSubmatch(plainText, -1)
	fmt.Println(plainTextResult)

	// match POSIX
	// POSIX 的模式是 [[ 开头 ]] 结束，两对方括号
	plainText = `body{
                   background-color: #fefbd8;
                 }
                 h1{
                   background-color: #0000ff;
                 }
                 div{
                   background-color: #d0f4e6;
                 }
                 span{
                   background-color: #f08970;
                 }`
	plainTextRegexp = regexp.MustCompile("[[:xdigit:]][[:xdigit:]][[:xdigit:]][[:xdigit:]][[:xdigit:]][[:xdigit:]]")
	if plainTextRegexp == nil {
		panic("plainTextRegexp err")
	}
	fmt.Println(".......match POSIX.............")
	plainTextResult = plainTextRegexp.FindAllStringSubmatch(plainText, -1)
	fmt.Println(plainTextResult)
}
