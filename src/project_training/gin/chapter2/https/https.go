package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
	"net/http"
)

type HttpsRes struct {
	Code   int    `json:"code"`
	Result string `json:"result"`
}

func main() {
	r := gin.Default()
	r.Use(httpsHandler())
	r.GET("/https", func(context *gin.Context) {
		fmt.Println(context.Request.Host)
		context.JSON(http.StatusOK, HttpsRes{
			Code:   http.StatusOK,
			Result: "succeed",
		})
	})
	path := "./src/chapter2/https/ca/"
	r.RunTLS(":9090", path+"ca.crt", path+"ca.key")
}

// go get github.com/unrolled/secure
func httpsHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		secureMiddle := secure.New(secure.Options{
			SSLRedirect: true, // 只允许 https 请求
			STSSeconds:  1536000,
		})
		err := secureMiddle.Process(context.Writer, context.Request)

		if err != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, "data is not safety")
			return
		}

		if status := context.Writer.Status(); status > 300 && status < 399 {
			context.Abort()
			return
		}
		context.Next()
	}
}
