package main

import (
	"context"
	"strconv"

	// 引入 proto 编译生成的包
	pb "github.com/laixhe/go-grpc/simple"
)

// 定义 User 并实现约定的接口
type User struct {
	pb.UnimplementedUserServer
	UserId   int64  `json:"user_id"`
	UserName string `json:"user_name"`
}

// 获取某个 user 数据
func (this *User) GetUser(ctx context.Context, ut *pb.GetUserRequest) (*pb.GetUserResponse, error) {

	// 待返回数据结构
	resp := new(pb.GetUserResponse)

	resp.Userid = ut.Userid
	resp.Username = "laixhe"
	resp.Sex = pb.UserSex_MEN

	return resp, nil
}

// 获取 user 所有数据
func (this *User) GetUserList(ctx context.Context, ut *pb.GetUserListRequest) (*pb.UserListResponse, error) {

	list := make([]*pb.GetUserResponse, 0, 3)

	for i := 1; i <= 3; i++ {

		list = append(list, &pb.GetUserResponse{Userid: int64(i), Username: "laiki" + strconv.Itoa(i), Sex: pb.UserSex_MEN})

	}

	// 待返回数据结构
	resp := new(pb.UserListResponse)
	resp.List = list

	return resp, nil
}
