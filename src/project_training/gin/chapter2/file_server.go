package main

import "github.com/gin-gonic/gin"

// 文件服务器
func main() {
	r := gin.Default()
	r.GET("/file", fileServer)
	r.Run(":9090")
}

func fileServer(context *gin.Context) {
	path := "./src/"
	filename := path + context.Query("name")
	context.File(filename)
}
