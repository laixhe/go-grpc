package main

import (
	"log"
	"net/http"

	// 引入 proto 编译生成的包
	pb "github.com/laixhe/go-grpc/simple"

	"google.golang.org/grpc"
)

const (
	// Address gRPC 服务地址
	Address = "127.0.0.1:50051"
)

var UClient pb.UserClient

// 初始化 Grpc 客户端
func initGrpc() {

	// 连接 GRPC 服务端
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	// 关闭连接
	//defer conn.Close()

	// 初始化 User 客户端
	UClient = pb.NewUserClient(conn)

	log.Println("初始化 Grpc 客户端成功")
	//select {}
}

// 启动 http 服务
func main() {

	initGrpc()

	http.HandleFunc("/user/get", GetUser)
	http.HandleFunc("/user/list", GetUserList)

	log.Println("开始监听 http 端口 0.0.0.0:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Printf("http.ListenAndServe err:%v", err)
	}
}
