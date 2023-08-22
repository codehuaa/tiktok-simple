/**
 * @Author: Keven5
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2023/8/22 17:50
 */

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"tiktok-simple/http/handlers"
	"tiktok-simple/http/rpc"
)

func Init() {
	rpc.InitRPC()
}

func main() {
	Init()

	h := server.New(
		server.WithHostPorts("0.0.0.0:8080"),
		server.WithHandleMethodNotAllowed(true),
	)

	// TODO：添加Jwt、log中间件

	// api 路由定义
	dyRouter := h.Group("/douyin")
	{
		// 用户相关api
		userRouter := dyRouter.Group("/user")
		{
			userRouter.POST("/register", handlers.Register)
			userRouter.POST("/login", handlers.Login)
			userRouter.GET("/", handlers.UserInfo)
		}
		// 视频相关api
		dyRouter.GET("/feed", handlers.Feed)
		videoRouter := dyRouter.Group("/publish")
		{
			videoRouter.POST("/action", handlers.PublishAction)
			videoRouter.GET("/list", handlers.PublishList)
		}
		// 赞相关api
		likeRouter := dyRouter.Group("/favorite")
		{
			likeRouter.POST("/action", handlers.FavoriteAction)
			likeRouter.GET("/list", handlers.FavoriteList)
		}
		// 评论相关api
		commentRouter := dyRouter.Group("/comment")
		{
			commentRouter.POST("/action", handlers.CommentAction)
			commentRouter.GET("/list", handlers.CommentList)
		}
	}

}
