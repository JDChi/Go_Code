package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"net/http"
)

// errgroup 是一个并发编程包，对 sync.WaitGroup 进行了封装，可以捕获到其中一个协程如果运行时的错误
// go get golang.org/x/sync/errgroup
var group errgroup.Group

// 多服务器运行
func main() {
	server1 := &http.Server{
		Addr:    ":9091",
		Handler: router1(),
	}
	server2 := &http.Server{
		Addr:    ":9092",
		Handler: router2(),
	}
	// 调用 Go 方法来开启一个线程进行运行
	group.Go(func() error {
		return server1.ListenAndServe()
	})

	group.Go(func() error {
		return server2.ListenAndServe()
	})

	if err := group.Wait(); err != nil {
		fmt.Println("execute failed : ", err)
	}
}

func router2() http.Handler {
	r1 := gin.Default()
	r1.GET("/server2", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "sever2 is response",
		})
	})
	return r1
}

func router1() http.Handler {
	r1 := gin.Default()
	r1.GET("/server1", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "sever1 is response",
		})
	})
	return r1
}
