package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

var log = logrus.New()

func initLogrus() error {
	log.Formatter = &logrus.JSONFormatter{}
	file, err := os.OpenFile("./gin_log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open file error")
		return err
	}

	log.Out = file
	// 将 gin 的日志输出交由 logrous 来处理，
	gin.DefaultWriter = log.Out
	log.Level = logrus.InfoLevel
	return nil
}

func main() {

	err := initLogrus()
	if err != nil {
		fmt.Println(err)
		return
	}

	r := gin.Default()
	r.GET("/logrus", func(context *gin.Context) {
		// 打印日志
		log.WithFields(logrus.Fields{
			"url":    context.Request.RequestURI,
			"method": context.Request.Method,
			"params": context.Query("name"),
			"IP":     context.ClientIP(),
		}).Info()
		resData := struct {
			Code int    `json:"code"`
			Msg  string `json:"msg"`
			Data string `json:"data"`
		}{http.StatusOK, "response succeed", "OK"}
		context.JSON(http.StatusOK, resData)
	})
	r.Run(":9090")
}
