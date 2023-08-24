/**
 * @Author: Keven5
 * @Description:
 * @File:  video
 * @Version: 1.0.0
 * @Date: 2023/8/22 16:57
 */

package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"tiktok-simple/kitex/kitex_gen/video"
	"tiktok-simple/kitex/kitex_gen/video/videoservice"
	"tiktok-simple/pkg/constants"
	"time"
)

var videoCLient videoservice.Client

func initVideoRpc() {

	c, err := videoservice.NewClient(
		constants.VideoServiceName,
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
	)
	if err != nil {
		panic(err)
	}
	videoCLient = c
}

// 下面的函数对应微服务中的服务，通过client调用微服务的服务
func Feed(ctx context.Context, req *video.FeedRequest) (resp *video.FeedResponse, err error) {
	return videoCLient.Feed(ctx, req)
}

func PublishAction(ctx context.Context, req *video.PublishActionRequest) (resp *video.PublishActionResponse, err error) {
	return videoCLient.PublishAction(ctx, req)
}

func PublishList(ctx context.Context, req *video.PublishListRequest) (resp *video.PublishListResponse, err error) {
	return videoCLient.PublishList(ctx, req)
}
