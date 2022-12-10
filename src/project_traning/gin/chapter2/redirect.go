package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 重定向
func main() {
	router := gin.Default()
	// 重定向到外部网络
	router.GET("/redirect", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	})

	// 重定向到具体路由
	router.GET("redirect2", func(context *gin.Context) {
		context.Request.URL.Path = "/testRedirect"
		router.HandleContext(context)
	})
	router.GET("/testRedirect", func(context *gin.Context) {
		context.String(http.StatusOK, "this is testRedirect")
	})
	router.Run("127.0.0.1:9090")
}
