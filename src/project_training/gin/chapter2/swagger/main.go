package main

import (
	_ "Go_Code/src/project_training/gin/chapter2/swagger/docs" // 需要添加 docs 路径引用
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

type User struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

// swagger
// 编写完注解后，使用 swag init 生成 docs 文件夹，里面包含相关生成文件
func main() {
	r := gin.Default()
	r.GET("/login", login)
	r.POST("/register", register)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":9090")
}

// @Tags register interface
// @Summary register
// @Accept json
// @Produce json
// @Router /register [post]
func register(context *gin.Context) {
	var user User
	err := context.Bind(&user)
	if err != nil {
		fmt.Println("bind error")
		context.JSON(http.StatusBadRequest, "bind error")
		return
	}
	res := Response{
		Code: http.StatusOK,
		Msg:  "register succeed",
		Data: "OK",
	}
	context.JSON(http.StatusOK, res)
}

// @Tags login interface
// @Summary login
// @Accept json
// @Produce json
// @Router /login [get]
func login(context *gin.Context) {
	username := context.Query("name")
	pwd := context.Query("pwd")
	fmt.Println(username, pwd)
	res := Response{}
	res.Code = http.StatusOK
	res.Msg = "login succeed"
	res.Data = "Ok"
	context.JSON(http.StatusOK, res)
}
