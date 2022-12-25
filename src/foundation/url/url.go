package main

import (
	"fmt"
	"net/url"
)

func main() {
	// 解析url
	urlParse, err := url.Parse("ttchat://channel_live?display_id=0&tabId=6")
	if err != nil {
		panic(err)
	}
	// 获取查询参数，得到的是个 map
	values := urlParse.Query()
	if _, ok := values["display_id"]; ok {
		// 更改参数值
		values.Set("display_id", "900")
		// 替换 url 的请求参数
		urlParse.RawQuery = values.Encode()
	}

	// 输出修改后的结果
	fmt.Printf("result %s", urlParse)
}
