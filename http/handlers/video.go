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
)

/**
 * Feed 处理视频流接口
 * 不限制登录状态，返回按投稿时间倒序的视频列表，视频数由服务端控制，单次最多30个。
 * @Author:
 */
func Feed(ctx context.Context, req *app.RequestContext) {

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
