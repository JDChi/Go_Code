package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 中间件
func main() {
	r := gin.Default()
	r.Use(MyMiddleware())
	r.GET("/middleware", func(context *gin.Context) {
		name := context.Query("name")
		ageStr := context.Query("age")
		ageInt, _ := strconv.Atoi(ageStr)
		res := struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}{name, ageInt}
		context.JSON(http.StatusOK, res)
	})
	r.Run(":9090")
}

func MyMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		name := context.Query("name")
		ageStr := context.Query("age")
		ageInt, err := strconv.Atoi(ageStr)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, "input wrong age, it is not integer")
			return
		}
		if ageInt < 0 || ageInt > 100 {
			context.AbortWithStatusJSON(http.StatusBadRequest, "input wrong age, it is not in range")
			return
		}
		if len(name) < 6 || len(name) > 12 {
			context.AbortWithStatusJSON(http.StatusBadRequest, "input wrong name")
			return
		}
		// 继续往下执行
		context.Next()
	}
}
