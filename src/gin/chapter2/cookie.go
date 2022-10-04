package main

import (
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 定义一个全局变量来存储相关的值
var cookieName string
var cookieValue string

// cookie 是保存在客户端本地，存在安全问题
func main() {
	r := gin.Default()
	r.Use(CookieMiddleware())
	r.GET("/cookie", func(context *gin.Context) {
		name := context.Query("name")
		if len(name) <= 0 {
			context.JSON(http.StatusBadRequest, "data wrong")
			return
		}
		cookieName = "cookie_" + name
		cookieValue = hex.EncodeToString([]byte(cookieName + "value"))
		val, _ := context.Cookie(cookieName)
		if val == "" {
			context.String(http.StatusOK, "You have got the cookie")
			return
		}
		context.String(http.StatusOK, "auth succeed, the cookie value is %s", val)

	})
	r.Run(":9090")

}

// CookieMiddleware 通过中间件来处理 cookie
func CookieMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		val, _ := context.Cookie(cookieName)
		if val == "" {
			context.SetCookie(cookieName, cookieValue, 3600, "/", "localhost", true, true)
			fmt.Println("The cookie has saved")
		}
	}
}
