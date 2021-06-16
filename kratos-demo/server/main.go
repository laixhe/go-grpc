package main

import (
	"log"

	"github.com/go-kratos/etcd/registry"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	etcdv3 "go.etcd.io/etcd/client/v3"

	// 引入 proto 编译生成的包
	pb "github.com/laixhe/go-grpc/goproto"
)

const (
	// ServeAddress 服务监听地址
	ServeAddress string = ":5501"
	// ServeName 服务名称
	ServeName string = "kratos-demo"
)

func main() {
	client, err := etcdv3.New(etcdv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
	})
	if err != nil {
		log.Fatal(err)
	}
	r := registry.New(client)

	grpcSrv := grpc.NewServer(
		grpc.Address(ServeAddress),
	)

	// 注册 User service
	pb.RegisterUserServer(grpcSrv, &User{})

	app := kratos.New(
		kratos.Name(ServeName),
		kratos.Version("latest"),
		kratos.Server(grpcSrv),
		kratos.Registrar(r),
	)

	err = app.Run()
	if err != nil {
		log.Fatal(err)
	}
}
