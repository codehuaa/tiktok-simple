package main

import (
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
	video "tiktok-simple/kitex/kitex_gen/video/videoservice"
	"tiktok-simple/service/video/service"
)

func main() {
	// 处理多个微服务时端口冲突的问题
	addr, _ := net.ResolveTCPAddr("tcp", ":8082")
	var opts []server.Option
	opts = append(opts, server.WithServiceAddr(addr))

	svr := video.NewServer(new(service.VideoServiceImpl), opts...)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
