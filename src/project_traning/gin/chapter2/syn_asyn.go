package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// 异步
func main() {
	r := gin.Default()
	r.GET("/async", func(context *gin.Context) {
		for i := 0; i < 6; i++ {
			cCp := context.Copy()
			go async(cCp, i)
			context.JSON(http.StatusOK, "async has run")
		}
	})
	r.Run(":9090")
}

func async(cp *gin.Context, i int) {
	fmt.Println("the no." + strconv.Itoa(i) + "goroutine is running")
	time.Sleep(4)
	fmt.Println("the no." + strconv.Itoa(i) + "goroutine is ended")
}
