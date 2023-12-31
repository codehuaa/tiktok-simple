/**
 * @Author: Keven5
 * @Description:
 * @File:  like
 * @Version: 1.0.0
 * @Date: 2023/8/22 16:57
 */

package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"tiktok-simple/kitex/kitex_gen/favorite"
	"tiktok-simple/kitex/kitex_gen/favorite/favoriteservice"
	"tiktok-simple/pkg/constants"
)

var likeClient favoriteservice.Client

func initLikeRpc() {

	c, err := favoriteservice.NewClient(
		constants.LikeServiceName,
		client.WithHostPorts("0.0.0.0:8083"),
		// client.WithMuxConnection(1),                    // mux
		// client.WithRPCTimeout(3*time.Second),           // rpc timeout
		// client.WithConnectTimeout(50*time.Millisecond), // conn timeout
	)
	if err != nil {
		panic(err)
	}
	likeClient = c
}

// 下面的函数对应微服务中的服务，通过client调用微服务的服务
func FavoriteAction(ctx context.Context, req *favorite.FavoriteActionRequest) (resp *favorite.FavoriteActionResponse, err error) {
	return likeClient.FavoriteAction(ctx, req)
}

func FavoriteList(ctx context.Context, req *favorite.FavoriteListRequest) (resp *favorite.FavoriteListResponse, err error) {
	return likeClient.FavoriteList(ctx, req)
}
