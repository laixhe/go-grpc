package main

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/laixhe/go-grpc/simple"
)

// 进行编解码 protobuf
func main() {

	res := new(simple.UserListResponse)
	var getdata []*simple.GetUserResponse

	for i := 0; i <= 3; i++ {

		getdata = append(getdata, &simple.GetUserResponse{Userid: int64(i), Username: "laiki", Sex: simple.UserSex_MEN})

	}
	res.List = getdata

	fmt.Println(res)

	// 进行编码 protobuf
	data, err := proto.Marshal(res)
	if err != nil {
		fmt.Println("res=", err)
		return
	}

	resp := new(simple.UserListResponse)

	// 进行解码 protobuf
	err = proto.Unmarshal(data, resp)
	if err != nil {
		fmt.Println("resp=", err)
		return
	}

	fmt.Println(resp)
}
