package main

import (
	helloPb "Go_Code/src/grpc/protopb/hello"
	"context"
	"google.golang.org/grpc"
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
	return &helloPb.HelloResp{
		Msg: "hello " + req.GetName()}, nil
}
