/**
 * @Author: Keven5
 * @Description:
 * @File:  comment
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
	"tiktok-simple/kitex/kitex_gen/comment"
	"tiktok-simple/pkg/response"
)

/**
 * CommentAction 评论操作
 * 登录用户对视频进行评论。
 * @Author:
 */
func CommentAction(ctx context.Context, req *app.RequestContext) {
	token := req.PostForm("token")
	vid, err := strconv.ParseInt(req.PostForm("video_id"), 10, 64)
	if err != nil {
		req.JSON(http.StatusOK, response.CommentAction{
			Base: response.Base{
				StatusCode: -1,
				StatusMsg:  "video_id 不合法",
			},
			Comment: nil,
		})
		return
	}
	actionType, err := strconv.ParseInt(req.PostForm("action_type"), 10, 64)
	if err != nil || (actionType != 1 && actionType != 2) {
		req.JSON(http.StatusOK, response.CommentAction{
			Base: response.Base{
				StatusCode: -1,
				StatusMsg:  "action_type 不合法",
			},
			Comment: nil,
		})
		return
	}
	requ := new(comment.CommentActionRequest)
	requ.Token = token
	requ.VideoId = vid
	requ.ActionType = int32(actionType)

	if actionType == 1 {
		commentText := req.PostForm("comment_text")
		if commentText == "" {
			req.JSON(http.StatusOK, response.CommentAction{
				Base: response.Base{
					StatusCode: -1,
					StatusMsg:  "comment_text 不能为空",
				},
				Comment: nil,
			})
			return
		}
		requ.CommentText = commentText
	} else if actionType == 2 {
		commentID, err := strconv.ParseInt(req.Query("comment_id"), 10, 64)
		if err != nil {
			req.JSON(http.StatusOK, response.CommentAction{
				Base: response.Base{
					StatusCode: -1,
					StatusMsg:  "comment_id 不合法",
				},
				Comment: nil,
			})
			return
		}
		requ.CommentId = commentID
	}
	res, _ := rpc.CommentAction(ctx, requ)
	if res.StatusCode == -1 {
		req.JSON(http.StatusOK, response.CommentAction{
			Base: response.Base{
				StatusCode: -1,
				StatusMsg:  res.StatusMsg,
			},
			Comment: nil,
		})
		return
	}
	req.JSON(http.StatusOK, response.CommentAction{
		Base: response.Base{
			StatusCode: 0,
			StatusMsg:  res.StatusMsg,
		},
		Comment: res.Comment,
	})
}

/**
 * CommentAction 视频评论列表
 * 查看视频的所有评论，按发布时间倒序。
 * @Author:
 */
func CommentList(ctx context.Context, req *app.RequestContext) {
	token := req.PostForm("token")
	vid, err := strconv.ParseInt(req.PostForm("video_id"), 10, 64)
	if err != nil {
		req.JSON(http.StatusOK, response.CommentList{
			Base: response.Base{
				StatusCode: -1,
				StatusMsg:  "video_id 不合法",
			},
			CommentList: nil,
		})
		return
	}
	r := &comment.CommentListRequest{
		Token:   token,
		VideoId: vid,
	}
	res, _ := rpc.CommentList(ctx, r)
	if res.StatusCode == -1 {
		req.JSON(http.StatusOK, response.CommentList{
			Base: response.Base{
				StatusCode: -1,
				StatusMsg:  res.StatusMsg,
			},
			CommentList: nil,
		})
		return
	}
	req.JSON(http.StatusOK, response.CommentList{
		Base: response.Base{
			StatusCode: 0,
			StatusMsg:  res.StatusMsg,
		},
		CommentList: res.CommentList,
	})
}
