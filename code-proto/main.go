package main

import (
	"fmt"

	"google.golang.org/protobuf/proto"

	// 引入 proto 编译生成的包
	pb "github.com/laixhe/go-grpc/goproto"
)

// 进行编解码 protobuf
func main() {

	res := new(pb.UserListResponse)
	var getdata []*pb.GetUserResponse

	for i := 0; i <= 3; i++ {

		getdata = append(getdata, &pb.GetUserResponse{Userid: int64(i), Username: "laiki", Sex: pb.UserSex_MEN})

	}
	res.List = getdata

	fmt.Println(res)

	// 进行编码 protobuf
	data, err := proto.Marshal(res)
	if err != nil {
		fmt.Println("res=", err)
		return
	}

	resp := new(pb.UserListResponse)

	// 进行解码 protobuf
	err = proto.Unmarshal(data, resp)
	if err != nil {
		fmt.Println("resp=", err)
		return
	}

	fmt.Println(resp)
}
