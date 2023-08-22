package main

import (
	"log"
	user "tiktok-simple/kitex/kitex_gen/user/userservice"
	user2 "tiktok-simple/service/user/service"
)

func main() {
	svr := user.NewServer(new(user2.UserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
