package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"path"
)

var (
	logFilePath = "./"
	logFileName = "system.log"
)

// 对日志进行最大时间设置和切割
func main() {
	r := gin.Default()
	r.Use(logMiddle())
	r.GET("/logrus2", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "succeed",
			"data": "OK",
		})

	})
	r.Run(":9090")
}

func logMiddle() gin.HandlerFunc {
	fileName := path.Join(logFilePath, logFileName)
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
	}
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	logger.Out = file

	return func(context *gin.Context) {

	}
}
