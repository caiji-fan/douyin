package redisutil

import (
	"douyin/config"

	"github.com/go-redis/redis/v8"
)

var RedisDB *redis.Client

// Init redis初始化
func Init() {
	// 建立连接
	opt, err := redis.ParseURL(config.Config.Redis.Url)
	if err != nil {
		panic(err)
	}

	RedisDB = redis.NewClient(opt)
	defer RedisDB.Close()
}
