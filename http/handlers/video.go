/**
 * @Author: Keven5
 * @Description:
 * @File:  video
 * @Version: 1.0.0
 * @Date: 2023/8/22 16:57
 */

package handlers

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"strconv"
	"tiktok-simple/http/rpc"
	"tiktok-simple/kitex/kitex_gen/video"
	"tiktok-simple/pkg/response"
	"time"
)

/**
 * Feed 处理视频流接口
 * 不限制登录状态，返回按投稿时间倒序的视频列表，视频数由服务端控制，单次最多30个。
 * @Author:
 */
func Feed(ctx context.Context, req *app.RequestContext) {
	token := req.PostForm("token")
	latestTime := req.PostForm("latest_time")
	var timestamp int64 = 0
	if latestTime != "" {
		timestamp, _ = strconv.ParseInt(latestTime, 10, 64)
	} else {
		timestamp = time.Now().UnixMilli()
	}

	r := &video.FeedRequest{
		LatestTime: timestamp,
		Token:      token,
	}
	res, _ := rpc.Feed(ctx, r)
	if res.StatusCode == -1 {
		req.JSON(http.StatusOK, response.Feed{
			Base: response.Err(res.StatusMsg),
		})
		return
	}
	req.JSON(http.StatusOK, response.Feed{
		Base:      response.OK,
		VideoList: res.VideoList,
	})
}

/**
 * PublishAction 处理视频投稿
 * 登录用户选择视频上传
 * @Author:
 */
func PublishAction(ctx context.Context, req *app.RequestContext) {

}

/**
 * PublishList 处理发布列表
 * 登录用户的视频发布列表，直接列出用户所有投稿过的视频。
 * @Author:
 */
func PublishList(ctx context.Context, req *app.RequestContext) {

}
