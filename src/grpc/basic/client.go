package main

import (
	helloPb "Go_Code/src/grpc/protopb/hello"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ClientTokenAuth struct {
}

// GetRequestMetadata 实现 grpc 的 credentials 包下的两个 token 认证接口
func (auth *ClientTokenAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appId":  "hello",
		"appKey": "123123",
	}, nil
}

func (auth *ClientTokenAuth) RequireTransportSecurity() bool {
	return false
}

func main() {

	var opts []grpc.DialOption
	// 选择禁用安全传输，无加密和验证
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithPerRPCCredentials(new(ClientTokenAuth)))

	conn, err := grpc.Dial("127.0.0.1:8028", opts...)
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
