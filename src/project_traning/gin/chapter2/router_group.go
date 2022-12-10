package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RespGroup struct {
	Data string
	Path string
}

// 路由分组
func main() {
	router := gin.Default()
	v1 := router.Group("v1")
	{
		r := v1.Group("/user")
		r.GET("/login", login)        // v1/user/login
		r2 := r.Group("/showInfo")    //v1/user/showInfo
		r2.GET("/abstract", abstract) // v1/user/showInfo/abstract
	}

	v2 := router.Group("v2")
	{
		v2.GET("/other", other) // v2/other
	}
	router.Run(":9091")

}

func other(context *gin.Context) {
	context.JSON(http.StatusOK, RespGroup{Data: "other", Path: context.Request.URL.Path})
}

func abstract(context *gin.Context) {
	context.JSON(http.StatusOK, RespGroup{Data: "abstract", Path: context.Request.URL.Path})

}

func login(context *gin.Context) {
	context.JSON(http.StatusOK, RespGroup{Data: "login", Path: context.Request.URL.Path})

}
