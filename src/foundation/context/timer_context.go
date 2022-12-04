package main

import (
	"context"
	"fmt"
	"time"
)

func timerContextHandleRequest(ctx context.Context) {
	go timerContextWriteRedis(ctx)
	go timerContextWriteDatabase(ctx)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("timerContextHandleRequest done")
			return
		default:
			fmt.Println("timerContextHandleRequest running")
			time.Sleep(2 * time.Second)
		}
	}
}

func timerContextWriteDatabase(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("timerContextWriteDatabase Done")
			return
		default:
			fmt.Println("timerContextWriteDatabase running")
			time.Sleep(2 * time.Second)
		}
	}
}

func timerContextWriteRedis(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("timerContextWriteRedis Done")
			return
		default:
			fmt.Println("timerContextWriteRedis running")
			time.Sleep(2 * time.Second)
		}
	}
}

// TimerContext give an example to set timeout to cancel all goroutines
func TimerContext() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	go timerContextHandleRequest(ctx)

	time.Sleep(10 * time.Second)

}
