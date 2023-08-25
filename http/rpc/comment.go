/**
 * @Author: Keven5
 * @Description:
 * @File:  comment
 * @Version: 1.0.0
 * @Date: 2023/8/22 16:57
 */

package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"tiktok-simple/kitex/kitex_gen/comment"
	"tiktok-simple/kitex/kitex_gen/comment/commentservice"
	"tiktok-simple/pkg/constants"
)

var commentClient commentservice.Client

func initCommentRpc() {

	c, err := commentservice.NewClient(
		constants.CommentServiceName,
		client.WithHostPorts("0.0.0.0:8084"),
		// client.WithMuxConnection(1),                    // mux
		// client.WithRPCTimeout(3*time.Second),           // rpc timeout
		// client.WithConnectTimeout(50*time.Millisecond), // conn timeout
	)
	if err != nil {
		panic(err)
	}
	commentClient = c
}

// 下面的函数对应微服务中的服务，通过client调用微服务的服务
func CommentAction(ctx context.Context, req *comment.CommentActionRequest) (resp *comment.CommentActionResponse, err error) {
	return commentClient.CommentAction(ctx, req)
}

func CommentList(ctx context.Context, req *comment.CommentListRequest) (resp *comment.CommentListResponse, err error) {
	return commentClient.CommentList(ctx, req)
}
