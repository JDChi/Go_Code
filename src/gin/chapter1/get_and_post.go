package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/get", getMsg)
	r.POST("/post", postMsg)
	r.Run("127.0.0.1:9090")

}

func postMsg(context *gin.Context) {
	name := context.DefaultPostForm("name", "")
	context.String(http.StatusOK, "Hello , %s", name)
}

func getMsg(context *gin.Context) {
	name := context.Query("name")
	context.String(http.StatusOK, "Hello , %s", name)
}
