package main

import (
	"log"
	comment "tiktok-simple/kitex/kitex_gen/comment/commentservice"
	comment2 "tiktok-simple/service/comment/service"
)

func main() {
	svr := comment.NewServer(new(comment2.CommentServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
