package main

import (
	"gin_learning/src/restgate"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 接口认证
// go get -u github.com/pjebs/restgate
// 在 Postman 里请求时，在 Header 那里加上 X-Auth-Key 和 X-Auth-Secret 作为 key，以及对应的 Value 即可请求
func main() {
	r := gin.Default()
	r.Use(APIAuthMiddleware())
	r.GET("/api_auth", func(context *gin.Context) {
		resData := struct {
			Code int    `json:"code"`
			Msg  string `json:"msg"`
			Data any    `json:"data"`
		}{http.StatusOK, "auth succeed", "Ok"}
		context.JSON(http.StatusOK, resData)
	})
	r.Run(":9090")
}

func APIAuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		gate := restgate.New("X-Auth-Key", "X-Auth-Secret", restgate.Static, restgate.Config{
			Key:                []string{"admin"},
			Secret:             []string{"adminpw"},
			HTTPSProtectionOff: true,
		})
		nextCalled := false
		nextAdapter := func(http.ResponseWriter, *http.Request) {
			nextCalled = true
			context.Next()
		}
		gate.ServeHTTP(context.Writer, context.Request, nextAdapter)
		if nextCalled == false {
			context.AbortWithStatus(401)
		}

	}
}
