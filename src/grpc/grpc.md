# grpc
## 安装
```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

## 生成 proto
```shell
protoc --go_out=./protopb ./proto/hello/hello.proto
protoc --go-grpc_out=./protopb ./proto/hello/hello.proto
```