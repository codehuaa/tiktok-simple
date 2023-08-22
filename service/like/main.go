package main

import (
	"log"
	favorite "tiktok-simple/kitex/kitex_gen/favorite/favoriteservice"
	"tiktok-simple/service/like/service"
)

func main() {
	svr := favorite.NewServer(new(service.FavoriteServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
