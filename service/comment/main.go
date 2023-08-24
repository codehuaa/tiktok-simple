package main

import (
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
	comment "tiktok-simple/kitex/kitex_gen/comment/commentservice"
	comment2 "tiktok-simple/service/comment/service"
)

func main() {
	// 处理多个微服务时端口冲突的问题
	addr, _ := net.ResolveTCPAddr("tcp", ":8084")
	var opts []server.Option
	opts = append(opts, server.WithServiceAddr(addr))

	svr := comment.NewServer(new(comment2.CommentServiceImpl), opts...)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
