package main

import (
	"context"
	"log"
	"net/http"

	"github.com/go-kratos/etcd/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	//"github.com/go-kratos/kratos/v2/transport/http"
	etcdv3 "go.etcd.io/etcd/client/v3"

	// 引入 proto 编译生成的包
	pb "github.com/laixhe/go-grpc/goproto"
)

// ServeName 服务名称
const ServeName string = "kratos-demo"

var UClient pb.UserClient

// 初始化 Grpc 客户端
func initGrpc() {
	cli, err := etcdv3.New(etcdv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
	})
	if err != nil {
		panic(err)
	}

	r := registry.New(cli)
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///"+ServeName),
		grpc.WithDiscovery(r),
	)
	if err != nil {
		log.Fatal(err)
	}

	UClient = pb.NewUserClient(conn)
	log.Println("初始化 Grpc 客户端成功")
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