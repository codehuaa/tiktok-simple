/**
 * @Author: Keven5
 * @Description:
 * @File:  init
 * @Version: 1.0.0
 * @Date: 2023/8/30 16:11
 */

package redis

import (
	"context"
	"github.com/go-redis/redis"
)

var _redis_client *redis.Client

func init() {
	redis_client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_redis_client = redis_client
}

func GetRedisClietn(ctx context.Context) *redis.Client {
	return _redis_client
}
