package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 多文件上传
func main() {
	r := gin.Default()
	r.POST("/uploads", func(context *gin.Context) {
		form, err := context.MultipartForm()
		if err != nil {
			context.String(http.StatusBadRequest, "upload file failed")

		}
		// 获取到相同 key 值的多个文件
		files := form.File["file_key"]
		dst := "./src/files/"
		// 对其进行遍历
		for _, file := range files {
			context.SaveUploadedFile(file, dst+file.Filename)
		}
		context.String(http.StatusOK, fmt.Sprintf("%d files upload succeed", len(files)))
	})
	r.Run(":9090")
}
