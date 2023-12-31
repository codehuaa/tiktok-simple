/**
 * @Author: Keven5
 * @Description:
 * @File:  resultCode
 * @Version: 1.0.0
 * @Date: 2023/8/25 15:55
 */

package response

var (
	// =============请求通用结构=============
	OK  = response(0, "ok")
	ERR = response(-1, "")

	// =============自定义的错误信息=============

)

func Err(msg string) Base {
	return response(-1, msg)
}
