/**
 * @Author: Keven5
 * @Description:
 * @File:  user
 * @Version: 1.0.0
 * @Date: 2023/8/22 16:57
 */

package handlers

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"tiktok-simple/http/rpc"
	kitex "tiktok-simple/kitex/kitex_gen/user"
)

/**
 * Register 用户注册接口
 * 新用户注册时提供用户名，密码，昵称即可，用户名需要保证唯一。创建成功后返回用户 id 和权限token.
 * @Author:
 */
func Register(ctx context.Context, req *app.RequestContext) {
	fmt.Println("call sucessfully!")
	username := req.Query("username")
	password := req.Query("password")
	userReq := &kitex.UserRegisterRequest{
		Username: username,
		Password: password,
	}
	resp, _ := rpc.Register(ctx, userReq)
	fmt.Println(resp)
}

/**
 * Login 用户注册接口
 * 通过用户名和密码进行登录，登录成功后返回用户 id 和权限 token.
 * @Author:
 */
func Login(ctx context.Context, req *app.RequestContext) {

}

/**
 * UserInfo 用户信息
 * 获取登录用户的 id、昵称，如果实现社交部分的功能，还会返回关注数和粉丝数。
 * @Author:
 */
func UserInfo(ctx context.Context, req *app.RequestContext) {

}
