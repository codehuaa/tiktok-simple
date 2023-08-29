/**
 * @Author: Keven5
 * @Description:
 * @File:  like
 * @Version: 1.0.0
 * @Date: 2023/8/22 16:57
 */

package handlers

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"strconv"
	"tiktok-simple/http/rpc"
	"tiktok-simple/kitex/kitex_gen/favorite"
	"tiktok-simple/pkg/response"
)

/**
 * FavoriteAction 赞操作
 * 登录用户对视频的点赞和取消点赞操作。
 * @Author:
 */
func FavoriteAction(ctx context.Context, req *app.RequestContext) {
	token := req.PostForm("token")
	video_id, err := strconv.ParseInt(req.PostForm("video_id"), 10, 64)
	if err != nil {
		req.JSON(consts.StatusOK, response.FavoriteAction{response.ERR.WithMsg("video_id不正确")})
		return
	}
	action_type, err := strconv.ParseInt(req.PostForm("action_type"), 10, 16)
	if err != nil {
		req.JSON(consts.StatusOK, response.FavoriteAction{response.ERR.WithMsg("action_type不正确")})
		return
	}
	like_action_req := &favorite.FavoriteActionRequest{
		Token:      token,
		VideoId:    video_id,
		ActionType: int32(action_type),
	}

	// rpc 调用
	_, err = rpc.FavoriteAction(ctx, like_action_req)
	if err != nil {
		req.JSON(consts.StatusOK, response.FavoriteAction{response.ERR.WithMsg(err.Error())})
		return
	}
	req.JSON(consts.StatusOK, response.FavoriteAction{response.OK})
}

/**
 * FavoriteList 喜欢列表
 * 登录用户的所有点赞视频。
 * @Author:
 */
func FavoriteList(ctx context.Context, req *app.RequestContext) {
	user_id, err := strconv.ParseInt(req.PostForm("user_id"), 10, 64)
	if err != nil {
		req.JSON(consts.StatusOK, response.FavoriteList{
			response.ERR.WithMsg("user_id不正确"), nil,
		})
		return
	}
	token := req.PostForm("token")

	// rpc
	like_list_req := &favorite.FavoriteListRequest{
		Token:  token,
		UserId: user_id,
	}
	data, err := rpc.FavoriteList(ctx, like_list_req)
	if err != nil {
		req.JSON(consts.StatusOK, response.FavoriteList{
			response.ERR.WithMsg(err.Error()), nil,
		})
		return
	}
	req.JSON(consts.StatusOK, response.FavoriteList{
		response.OK, data.VideoList,
	})
}
