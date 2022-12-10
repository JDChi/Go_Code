package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 多形式渲染数据格式
func main() {
	r := gin.Default()
	// json 格式输出
	r.GET("/json", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"html": "<b>hello gin</b>",
		})
	})

	r.GET("/someHTML", func(context *gin.Context) {
		context.PureJSON(http.StatusOK, gin.H{
			"html": "<b>hello gin</b>",
		})
	})

	r.GET("/someXML", func(context *gin.Context) {
		type Message struct {
			Name string
			Msg  string
			Age  int
		}
		info := Message{}
		info.Name = "ddd"
		info.Msg = "world"
		info.Age = 20
		context.XML(http.StatusOK, info)
	})

	r.GET("someYAML", func(context *gin.Context) {
		context.YAML(http.StatusOK, gin.H{
			"message": "hello world",
			"status":  200,
		})
	})
	r.Run(":9090")
}
