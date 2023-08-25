/**
 * @Author: Keven5
 * @Description:
 * @File:  resposneType
 * @Version: 1.0.0
 * @Date: 2023/8/25 22:57
 */

package response

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
}

// 视频相关

// 喜欢相关

// 评论相关
