package main

import (
	"fmt"
	"regexp"
)

func main() {
	var plainText = `Send personal email to ben@forta.com or ben.forta@forta.com. 
                     For questions about a book use support@forta.com. If your message
                     is urgent try ben@urgent.forta.com.`
	// + 表示匹配一个或多个字符，也就是至少要有一个
	plainTextRegexp := regexp.MustCompile("\\w+@\\w+.\\w+")
	if plainTextRegexp == nil {
		panic("plainTextRegexp err")
	}
	fmt.Println(".......match one chars at least.............")
	// 匹配多个连续重复出现的字符，只能匹配非常常规的邮件地址
	plainTextResult := plainTextRegexp.FindAllStringSubmatch(plainText, -1)
	fmt.Println(plainTextResult)

	// 以 \w+ 结尾，这样匹配出来的邮件不会出现句号 . , 如果以 [\w.]+ 结尾，则会把句号也算进去
	plainTextRegexp = regexp.MustCompile("[\\w.]+@[\\w.]+\\.\\w+")
	if plainTextRegexp == nil {
		panic("plainTextRegexp err")
	}
	fmt.Println(".......match one chars2 at least.............")
	// 匹配多个连续重复出现的字符，进一步完善匹配条件
	plainTextResult = plainTextRegexp.FindAllStringSubmatch(plainText, -1)
	fmt.Println(plainTextResult)

	plainText = `Hello .ben@forta.com is my email address.`
	// * 表示 0 个或多个字符，相比于 +，它可以一个都没有
	plainTextRegexp = regexp.MustCompile("\\w[\\w+]*@[\\w.]+\\.\\w+")
	if plainTextRegexp == nil {
		panic("plainTextRegexp err")
	}
	fmt.Println(".......match zero or multiple chars.............")
	// 匹配多个连续重复出现的字符，进一步完善匹配条件
	plainTextResult = plainTextRegexp.FindAllStringSubmatch(plainText, -1)
	fmt.Println(plainTextResult)

	plainText = "The URL is http://www.forta.com/, to connect securely use https://www.forta.com/ instead"
	// ? 表示 0 个或 1 个字符，也就是最多出现一次
	plainTextRegexp = regexp.MustCompile("https?://[\\w./]+")
	if plainTextRegexp == nil {
		panic("plainTextRegexp err")
	}
	fmt.Println(".......match one char at most.............")
	// 匹配多个连续重复出现的字符，进一步完善匹配条件
	plainTextResult = plainTextRegexp.FindAllStringSubmatch(plainText, -1)
	fmt.Println(plainTextResult)

	// 上面的 +、*、？ 都不能指定重复的次数，可以使用 {} 来指定次数，即必须达到该值，小于也不算
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
	plainTextRegexp = regexp.MustCompile("[[:xdigit:]]{6}")
	if plainTextRegexp == nil {
		panic("plainTextRegexp err")
	}
	fmt.Println(".......match interval.............")
	plainTextResult = plainTextRegexp.FindAllStringSubmatch(plainText, -1)
	fmt.Println(plainTextResult)

	plainText = `4/8/17
                  10-6-2018
                  2/2/2
                  01-01-01`
	// {} 也可以指定重复次数的范围
	plainTextRegexp = regexp.MustCompile("\\d{1,2}[/-]\\d{1,2}[/-]\\d{2,4}")
	if plainTextRegexp == nil {
		panic("plainTextRegexp err")
	}
	fmt.Println(".......match interval range.............")
	plainTextResult = plainTextRegexp.FindAllStringSubmatch(plainText, -1)
	fmt.Println(plainTextResult)

	plainText = `1001: $496.80
                 1002: $1290.69
                 1003: $26.43
                 1004: $613.42
                 1005: $7.61
                 1006: $414.90
                 1007: $25.00`
	// 匹配出至少是 100 美元的账单 {3,}表示至少是 3 次，3 位数字则代表至少是 100
	plainTextRegexp = regexp.MustCompile("\\d+: \\$\\d{3,}\\.\\d+")
	if plainTextRegexp == nil {
		panic("plainTextRegexp err")
	}
	fmt.Println(".......match interval range at least.............")
	plainTextResult = plainTextRegexp.FindAllStringSubmatch(plainText, -1)
	fmt.Println(plainTextResult)

	// 防止过度匹配
	plainText = `This offer is not available to customers living in <b>AK</b> and <b>HI</b>`

	//plainTextRegexp = regexp.MustCompile("<[Bb]>.*</[Bb]>")
	//上面的结果是 <b>AK</b> and <b>HI</b>
	if plainTextRegexp == nil {
		panic("plainTextRegexp err")
	}
	fmt.Println(".......match interval range at least.............")
	plainTextResult = plainTextRegexp.FindAllStringSubmatch(plainText, -1)
	fmt.Println(plainTextResult)

}
