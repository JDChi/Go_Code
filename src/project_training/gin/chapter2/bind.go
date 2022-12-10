package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Remark   string
}

// 绑定 将 client 传递的 json 数据绑定到结构体上
func main() {
	r := gin.Default()
	r.POST("/login", func(context *gin.Context) {
		var login Login
		err := context.Bind(&login)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"msg":  "bind failed",
				"data": err.Error(),
			})
			return
		}
		if login.UserName == "user" && login.Password == "123456" {
			context.JSON(http.StatusOK, gin.H{
				"msg":  "login succeed",
				"data": "ok",
			})
			return
		}
		context.JSON(http.StatusBadRequest, gin.H{
			"msg":  "login failed",
			"data": "error",
		})
		return
	})
	r.Run(":9090")

}
