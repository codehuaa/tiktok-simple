/**
 * @Author: Keven5
 * @Description:
 * @File:  util
 * @Version: 1.0.0
 * @Date: 2023/8/22 17:10
 */

package constants

import "os"

func GetIp(key string) string {
	ip := os.Getenv(key)
	if ip == "" {
		ip = "localhost"
	}
	return ip
}
