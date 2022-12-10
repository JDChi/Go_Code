package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 登录验证中间件
func main() {
	r := gin.Default()
	r.Use(AuthMiddleware())
	r.GET("/login", func(context *gin.Context) {
		user := context.MustGet(gin.AuthUserKey).(string)
		context.JSON(http.StatusOK, "login succeed : "+user)
	})
	r.Run(":9090")
}

func AuthMiddleware() gin.HandlerFunc {
	accounts := gin.Accounts{
		"admin": "admin",
	}
	auth := gin.BasicAuth(accounts)
	return auth
}
