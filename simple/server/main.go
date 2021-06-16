package main

import (
	"log"
	"net"

	// 引入 proto 编译生成的包
	pb "github.com/laixhe/go-grpc/goproto"

	"google.golang.org/grpc"
)

func main() {

	// 监听地址和端口
	listen, err := net.Listen("tcp", ":5501")
	if err != nil {
		log.Fatalf("监听端口失败: %v", err)
	}

	// 实例化 grpc Server
	serverGrpc := grpc.NewServer()

	// 注册 User service
	pb.RegisterUserServer(serverGrpc, &User{})

	log.Println("开始监听 Grpc 端口 0.0.0.0:5501")

	// 启动服务
	err = serverGrpc.Serve(listen)
	if err != nil {
		log.Println("启动 Grpc 服务失败")
	}
}
