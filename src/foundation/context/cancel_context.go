package main

import (
	"context"
	"fmt"
	"time"
)

func cancelContextHandleRequest(ctx context.Context) {
	go cancelContextWriteRedis(ctx)
	go cancelContextWriteDatabase(ctx)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("cancelContextHandleRequest done")
			return
		default:
			fmt.Println("cancelContextHandleRequest running")
			time.Sleep(2 * time.Second)
		}
	}
}

func cancelContextWriteDatabase(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("cancelContextWriteDatabase Done")
			return
		default:
			fmt.Println("cancelContextWriteDatabase running")
			time.Sleep(2 * time.Second)
		}
	}
}

func cancelContextWriteRedis(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("cancelContextWriteRedis Done")
			return
		default:
			fmt.Println("cancelContextWriteRedis running")
			time.Sleep(2 * time.Second)
		}
	}
}

// CancelContext give an example to cancel all sub goroutine
func CancelContext() {
	ctx, cancel := context.WithCancel(context.Background())
	go cancelContextHandleRequest(ctx)

	time.Sleep(5 * time.Second)
	fmt.Println("It is time to stop all sub goroutines")
	cancel()

	time.Sleep(5 * time.Second)

}
