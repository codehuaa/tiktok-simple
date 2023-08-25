/**
 * @Author: Keven5
 * @Description: 封装http中的response
 * @File:  response.go
 * @Version: 1.0.0
 * @Date: 2023/8/25 15:48
 */

package response

// Status_Code 为0表示成功，其他值表示失败
type Base struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

// WithMsg 快速创建一个只更改 Msg 的Response
func (resp *Base) WithMsg(message string) Base {
	return Base{
		StatusCode: resp.StatusCode,
		StatusMsg:  message,
	}
}

// response 构造一个 Response 框架，不携带 Data
func response(code int, msg string) Base {
	return Base{
		StatusCode: code,
		StatusMsg:  msg,
	}
}
