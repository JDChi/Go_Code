package main

import (
	"context"
	"fmt"
	"time"
)

func valueContextHandleRequest(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("valueContextHandleRequest Done")
			return
		default:
			fmt.Println("valueContextHandleRequest running, param = ", ctx.Value("param"))
			time.Sleep(2 * time.Second)
		}
	}
}

// ValueContext give an example to pass a value to sub goroutine
func ValueContext() {
	ctx := context.WithValue(context.Background(), "param", "hello")
	go valueContextHandleRequest(ctx)

	time.Sleep(10 * time.Second)

}
