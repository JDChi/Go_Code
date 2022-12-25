package main

import (
	helloPb "Go_Code/src/grpc/protopb/hello"
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net"
)

type Server struct {
	helloPb.UnimplementedSayHelloServer
}

func main() {
	lis, err := net.Listen("tcp", ":8028")
	if err != nil {
		panic(err)
	}

	// 创建 grpc 服务
	grpcServer := grpc.NewServer()
	// 注册我们编写的服务
	helloPb.RegisterSayHelloServer(grpcServer, &Server{})
	// 启动服务
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}

}

func (s *Server) SayHello(ctx context.Context, req *helloPb.HelloReq) (*helloPb.HelloResp, error) {
	// 获取元数据的信息，实际是这部分 token 的东西写在拦截器里
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("no token")
	}

	fmt.Printf("md = %v", md)
	var (
		appId, appKey string
	)

	if v, ok := md["appid"]; ok {
		appId = v[0]
	}

	if v, ok := md["appkey"]; ok {
		appKey = v[0]
	}
	fmt.Printf("appId = %v, appKey = %v\n", appId, appKey)
	if appId != "hello" || appKey != "123123" {
		return nil, errors.New("token wrong")
	}

	return &helloPb.HelloResp{
		Msg: "hello " + req.GetName()}, nil
}
