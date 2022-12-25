package main

import (
	helloPb "Go_Code/src/grpc/protopb/hello"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	// 选择禁用安全传输，无加密和验证
	conn, err := grpc.Dial("127.0.0.1:8028", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	// 记得要关闭
	defer conn.Close()

	helloClient := helloPb.NewSayHelloClient(conn)
	resp, err := helloClient.SayHello(context.Background(), &helloPb.HelloReq{
		Name: "Jack",
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("resp = %v", resp)

}
