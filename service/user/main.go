package main

import (
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
	user "tiktok-simple/kitex/kitex_gen/user/userservice"
	user2 "tiktok-simple/service/user/service"
)

func Init() {
	user2.Init("signingKey")
}

func main() {

	Init()

	// 处理多个微服务时端口冲突的问题
	addr, _ := net.ResolveTCPAddr("tcp", ":8081")
	var opts []server.Option
	opts = append(opts, server.WithServiceAddr(addr))

	svr := user.NewServer(new(user2.UserServiceImpl), opts...)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
