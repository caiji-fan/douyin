package redisutil

import (
	"context"
	"douyin/config"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var RedisDB *redis.Client

// Init redis初始化
func Init() {
	// 建立连接
	RedisDB = redis.NewClient(&redis.Options{
		Addr:     config.Config.Redis.Url,
		Password: "",
		DB:       0,
	})

	err := RedisDB.Ping(context.Background()).Err()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}
