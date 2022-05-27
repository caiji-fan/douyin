package redisutil

import (
	"douyin/config"
	"log"

	"github.com/go-redis/redis"
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
	err := RedisDB.Ping().Err()
	if err != nil {
		log.Fatalln(err)
	}
}
