/**
 * @Author: Keven5
 * @Description:
 * @File:  init
 * @Version: 1.0.0
 * @Date: 2023/8/22 16:57
 */

package rpc

// InitRPC 初始化所有问项目需要使用到的rpc-client
func InitRPC() {
	initUserRpc()
	initVideoRpc()
	initLikeRpc()
	initCommentRpc()
}
