rm *.go

protoc --go-grpc_out=. --go_out=. *.proto