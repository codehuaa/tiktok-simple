package main

import (
	"log"
	video "tiktok-simple/kitex/kitex_gen/video/videoservice"
	"tiktok-simple/service/video/service"
)

func main() {
	svr := video.NewServer(new(service.VideoServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
