package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 文件上传
func main() {
	r := gin.Default()
	r.POST("/upload", func(context *gin.Context) {
		file, err := context.FormFile("filename")
		if err != nil {
			context.String(http.StatusBadRequest, "upload file failed")
		}
		dst := "./src/files"
		context.SaveUploadedFile(file, dst+file.Filename)
		context.String(http.StatusOK, fmt.Sprintf("%s upload file succeed", file.Filename))
	})
	r.Run(":9090")

}
