// Package redisutil
// @Author shaofan
// @Date 2022/5/13
// @DESC redis连接初始化与工具
package redisutil

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

// Set 							插入string类型
// key 							键
// value 						需要插入的值，内部需要进行序列化
func Set[T any](key string, value *T) error {
	err := RedisDB.Set(context.Background(), key, value, -1).Err()
	return err
}

// SetWithExpireTime 			插入string类型并设置过期时间
// key 							键
// value 						需要插入的值，内部需要进行序列化
// duration						过期时间
func SetWithExpireTime[T any](key string, value *T, duration time.Duration) error {
	err := RedisDB.Set(context.Background(), key, value, duration).Err()
	return err
}

// Get 							获取string类型
// key 							键
// value 						获取的值存储的指针
func Get[T any](key string, value *T) error {
	v, err := RedisDB.Get(context.Background(), key).Result()
	if err == redis.Nil {
		fmt.Print("key dose not exist")
	} else if err != nil {
		fmt.Printf("get %v failed, err:%v\n", key, err)
		return err
	}
	fmt.Printf("get %v succeed, value:%v\n", key, v)
	return nil
}

// GetAndDelete					获取string并删除string类型
// key							键
// value						值
func GetAndDelete[T any](key string, value *T) error {
	v, err := RedisDB.Get(context.Background(), key).Result()
	if err == redis.Nil {
		fmt.Print("key dose not exist")
	} else if err != nil {
		fmt.Printf("get %v failed, err:%v\n", key, err)
	}
	fmt.Printf("get %v succeed, value:%v\n", key, v)

	_, err = RedisDB.Del(context.Background(), key).Result()
	if err == redis.Nil {
		fmt.Print("key dose not exist")
	} else if err != nil {
		fmt.Printf("del %v failed, err:%v\n", key, err)
	}
	fmt.Printf("del %v succeed", key)
	return nil
}

// ZSet 						插入set类型
// key						 	键
// value 						需要插入的值
// score 						排序字段名，如果没有该字段则不进行排序
func ZSet[T any](key string, value *[]T, score string) error {
	addValues := make([]*redis.Z, len(*value))
	if score != "" {
		scoref64, err := strconv.ParseFloat(score, 64)
		if err != nil {
		}
		for _, value := range *value {
			addValues = append(addValues, &redis.Z{
				Score:  scoref64,
				Member: value,
			})
		}
	} else {
		for _, value := range *value {
			addValues = append(addValues, &redis.Z{
				Member: value,
			})
		}
	}
	_, err := RedisDB.ZAdd(context.Background(), key, addValues...).Result()
	return err
}

// ZGet 						获取set类型
// key 							键
// value 						获取的值存储的指针
func ZGet(key string, value interface{}) error {
	score, err := RedisDB.ZRange(context.Background(), key, 0, -1).Result()
	fmt.Printf("zget 获得值：%v", score)
	value = &score
	return err
}

// ZSetWithExpireTime 			插入set类型并设置过期时间
// key							键
// value						需要插入的值
// score						排序字段名，如果没有该字段则不进行排序
// duration						时间
func ZSetWithExpireTime[T any](key string, value *[]T, score string, duration time.Duration) error {
	ok, _ := RedisDB.Expire(context.Background(), key, duration).Result()
	if ok {
		fmt.Println("name 过期时间设置成功", ok)
	} else {
		fmt.Println("name 过期时间设置失败", ok)
	}

	addValues := make([]*redis.Z, len(*value))
	if score != "" {
		scoref64, err := strconv.ParseFloat(score, 64)
		if err != nil {
		}
		for _, value := range *value {
			addValues = append(addValues, &redis.Z{
				Score:  scoref64,
				Member: value,
			})
		}
	} else {
		for _, value := range *value {
			addValues = append(addValues, &redis.Z{
				Member: value,
			})
		}
	}
	_, err := RedisDB.ZAdd(context.Background(), key, addValues...).Result()
	return err
}
