/**
 * @Author: Keven5
 * @Description:
 * @File:  resposneType
 * @Version: 1.0.0
 * @Date: 2023/8/25 22:57
 */

package response

import (
	"tiktok-simple/kitex/kitex_gen/comment"
	"tiktok-simple/kitex/kitex_gen/user"
	"tiktok-simple/kitex/kitex_gen/video"
)

// 用户相关
type Register struct {
	Base
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

type Login struct {
	Base
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

type UserInfo struct {
	Base
	user.User
}

// 视频相关
type PublishAction struct {
	Base
}

type PublishList struct {
	Base
	VideoList []*video.Video `json:"video_list"`
}

type Feed struct {
	Base
	NextTime  int64          `json:"next_time"`
	VideoList []*video.Video `json:"video_list"`
}

// 喜欢相关
type FavoriteAction struct {
	Base
}

type FavoriteList struct {
	Base
	VideoList []*video.Video `json:"video_list"`
}

// 评论相关
type CommentAction struct {
	Base
	Comment *comment.Comment `json:"comment"`
}

type CommentList struct {
	Base
	CommentList []*comment.Comment `json:"comment_list"`
}

type RelationAction struct {
	Base
}
