/**
 * @Author: Keven5
 * @Description:
 * @File:  init
 * @Version: 1.0.0
 * @Date: 2023/8/26 14:25
 */

package service

import "tiktok-simple/pkg/jwt"

var Jwt *jwt.JWT

func Init(signingKey string) {
	Jwt = jwt.NewJWT([]byte(signingKey))
}
