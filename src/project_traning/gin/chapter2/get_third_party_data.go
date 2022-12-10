package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 在 client 请求时，获取第三方网站上的数据，将其返回
func main() {
	router := gin.Default()
	router.GET("/GetOtherData", func(context *gin.Context) {
		url := "http://www.baidu.com"
		resp, err := http.Get(url)
		if err != nil || resp.StatusCode != http.StatusOK {
			context.Status(http.StatusServiceUnavailable)
			return
		}
		body := resp.Body
		contentLength := resp.ContentLength
		contentType := resp.Header.Get("Content-Type")
		context.DataFromReader(http.StatusOK, contentLength, contentType, body, nil)
	})
	router.Run(":9090")

}
