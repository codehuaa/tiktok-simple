/**
 * @Author: Keven5
 * @Description:
 * @File:  video
 * @Version: 1.0.0
 * @Date: 2023/8/22 16:57
 */

package handlers

import (
	"bytes"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"io"
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
	token := req.PostForm("token")
	if token == "" {
		req.JSON(http.StatusOK, response.PublishAction{
			Base: response.Err("用户鉴权失败，token为空"),
		})
		return
	}
	title := req.PostForm("title")
	if title == "" {
		req.JSON(http.StatusOK, response.PublishAction{
			Base: response.Err("标题不能为空"),
		})
		return
	}
	// 视频数据
	file, err := req.FormFile("data")
	if err != nil {
		req.JSON(http.StatusBadRequest, response.RelationAction{
			Base: response.Err("上传视频加载失败"),
		})
		return
	}
	src, err := file.Open()
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, src); err != nil {
		req.JSON(http.StatusBadRequest, response.RelationAction{
			Base: response.Err("视频上传失败"),
		})
		return
	}

	r := &video.PublishActionRequest{
		Token: token,
		Title: title,
		Data:  buf.Bytes(),
	}
	res, _ := rpc.PublishAction(ctx, r)
	if res.StatusCode == -1 {
		req.JSON(http.StatusOK, response.PublishAction{
			Base: response.Err(res.StatusMsg),
		})
		return
	}
	req.JSON(http.StatusOK, response.PublishAction{
		Base: response.OK,
	})
}

/**
 * PublishList 处理发布列表
 * 登录用户的视频发布列表，直接列出用户所有投稿过的视频。
 * @Author:
 */
func PublishList(ctx context.Context, req *app.RequestContext) {
	token := req.GetString("token")

	uid, err := strconv.ParseInt(req.PostForm("user_id"), 10, 64)
	if err != nil {
		req.JSON(http.StatusOK, response.PublishList{
			Base: response.Err("user_id不合法"),
		})
		return
	}
	r := &video.PublishListRequest{
		Token:  token,
		UserId: uid,
	}
	res, _ := rpc.PublishList(ctx, r)
	if res.StatusCode == -1 {
		req.JSON(http.StatusOK, response.PublishList{
			Base: response.Err(res.StatusMsg),
		})
		return
	}
	req.JSON(http.StatusOK, response.PublishList{
		Base:      response.OK,
		VideoList: res.VideoList,
	})
}
