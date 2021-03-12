#### 确保开启了 go module 模式
```
# 代理
GOPROXY=https://goproxy.io,direct
# 当前项目下
go mod tidy
```

#### 生成工具 - protoc 插件
```
# protoc v3.15.5
https://github.com/protocolbuffers/protobuf/releases

# protoc-gen-go v1.25.0
https://github.com/protocolbuffers/protobuf-go/releases

# protoc-gen-go-grpc v1.1.0
https://github.com/grpc/grpc-go/releases
```

#### proto 文件编译
```
protoc --go-grpc_out=. --go_out=. *.proto

protoc 参数说明
--proto_path     proto 文件目录(简写 -I 如果没有指定参数，则在当前目录进行搜索)
--go-grpc_out=   生成的 go-grpc 源码保存目录
--go_out         生成的 go 源码保存目录
```

#### etcd 单机版 v3.4.15
```
./etcd --listen-client-urls=http://0.0.0.0:2379 --advertise-client-urls=http://0.0.0.0:2379
```
