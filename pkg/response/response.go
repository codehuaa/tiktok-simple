/**
 * @Author: Keven5
 * @Description: 封装http中的response
 * @File:  response.go
 * @Version: 1.0.0
 * @Date: 2023/8/25 15:48
 */

package response

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// WithMsg 快速创建一个只更改 Msg 的Response
func (resp *Response) WithMsg(message string) Response {
	return Response{
		Code: resp.Code,
		Msg:  message,
		Data: resp.Data,
	}
}

// WithData 快速创建一个只更改 Data 的Response
func (resp *Response) WithData(data interface{}) Response {
	return Response{
		Code: resp.Code,
		Msg:  resp.Msg,
		Data: data,
	}
}

// response 构造一个 Response 框架，不携带 Data
func response(code int, msg string) *Response {
	return &Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}
