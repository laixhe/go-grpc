#### 需要 Golang 1.16 以上
> 确保开启了 go module 模式

```
# 设置代理
GOPROXY=https://goproxy.io,direct
# 当前项目下
go mod tidy
```

#### 生成工具 - protoc 插件
```
# protoc
https://github.com/protocolbuffers/protobuf/releases/tag/v3.15.8
protoc --version

# protoc-gen-go
go get google.golang.org/protobuf/cmd/protoc-gen-go@v1.26.0
protoc-gen-go --version

# protoc-gen-go-grpc
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0
protoc-gen-go-grpc --version
```

#### proto 文件编译
```
protoc --go-grpc_out=. --go_out=. *.proto

protoc 参数说明
--proto_path     proto 文件目录(简写 -I 如果没有指定参数，则在当前目录进行搜索)
--go-grpc_out=   生成的 go-grpc 源码保存目录
--go_out         生成的 go 源码保存目录
```

#### etcd 单机版 v3.4.16
```
./etcd --listen-client-urls=http://0.0.0.0:2379 --advertise-client-urls=http://0.0.0.0:2379
```