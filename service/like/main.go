package main

import (
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
	favorite "tiktok-simple/kitex/kitex_gen/favorite/favoriteservice"
	"tiktok-simple/service/like/service"
)

func Init() {
	service.Init("signingKey")
}

func main() {
	Init()

	// 处理多个微服务时端口冲突的问题
	addr, _ := net.ResolveTCPAddr("tcp", ":8083")
	var opts []server.Option
	opts = append(opts, server.WithServiceAddr(addr))

	svr := favorite.NewServer(new(service.FavoriteServiceImpl), opts...)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
