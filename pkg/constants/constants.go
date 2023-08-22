/**
 * @Author: Keven5
 * @Description:
 * @File:  constants
 * @Version: 1.0.0
 * @Date: 2023/8/22 17:08
 */

package constants

var (
	UserServiceName    = "userService"
	VideoServiceName   = "videoService"
	LikeServiceName    = "likeService"
	CommentServiceName = "commentService"
)

var (
	EtcdAddress = GetIp("EtcdIp") + ":2379"
)
