package main

import (
	"context"
	"log"
	"net"

	// 引入 proto 编译生成的包
	pb "github.com/laixhe/go-grpc/goproto"

	"google.golang.org/grpc"
)

// 一元拦截器
// 函数名无特殊要求，参数需一致
// req 包含请求的所有信息，info 包含一元RPC服务的所有信息
func unaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	// 前置处理逻辑
	log.Printf("[unary interceptor request] %s", info.FullMethod)
	// 完成RPC服务的正常执行
	m, err := handler(ctx, req)
	// 后置处理逻辑
	log.Printf("[unary interceptor resonse] %s", m)
	// 返回响应
	return m, err
}

func main() {

	// 监听地址和端口
	listen, err := net.Listen("tcp", ":5501")
	if err != nil {
		log.Fatalf("监听端口失败: %v", err)
	}

	// 实例化 grpc 服务器
	// 在创建 gRPC 服务器实例的时候注册拦截器
	// NewServer 可传入多个拦截器
	serverGrpc := grpc.NewServer(grpc.UnaryInterceptor(unaryServerInterceptor))

	// 将服务注册到 grpc 服务器上
	pb.RegisterUserServer(serverGrpc, &User{})

	log.Println("开始监听 Grpc 端口 0.0.0.0:5501")

	// 绑定 grpc 服务器到指定 tcp
	// 启动服务
	err = serverGrpc.Serve(listen)
	if err != nil {
		log.Println("启动 Grpc 服务失败")
	}
}
