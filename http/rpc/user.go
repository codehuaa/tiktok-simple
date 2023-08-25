/**
 * @Author: Keven5
 * @Description:
 * @File:  user
 * @Version: 1.0.0
 * @Date: 2023/8/22 16:57
 */

package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"tiktok-simple/kitex/kitex_gen/user"
	"tiktok-simple/kitex/kitex_gen/user/userservice"
	"tiktok-simple/pkg/constants"
)

var userClient userservice.Client

func initUserRpc() {
	c, err := userservice.NewClient(
		constants.UserServiceName,
		client.WithHostPorts("0.0.0.0:8081"),
		// client.WithMuxConnection(1),                    // mux
		// client.WithRPCTimeout(3*time.Second),           // rpc timeout
		// client.WithConnectTimeout(50*time.Millisecond), // conn timeout
	)
	if err != nil {
		panic(err)
	}
	userClient = c
}

// 下面的函数对应微服务中的服务，通过client调用微服务的服务
func Register(ctx context.Context, req *user.UserRegisterRequest) (resp *user.UserRegisterResponse, err error) {
	return userClient.Register(ctx, req)
}

func Login(ctx context.Context, req *user.UserLoginRequest) (resp *user.UserLoginResponse, err error) {
	return userClient.Login(ctx, req)
}

func UserInfo(ctx context.Context, req *user.UserInfoRequest) (resp *user.UserInfoResponse, err error) {
	return userClient.UserInfo(ctx, req)
}
