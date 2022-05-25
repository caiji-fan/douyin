// Package redisutil
// @Author shaofan
// @Date 2022/5/13
// @DESC redis连接初始化与工具
package redisutil

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

// Set 							插入string类型
// key 							键
// value 						需要插入的值，内部需要进行序列化
func Set(key string, value interface{}) error {
	err := RedisDB.Set(context.Background(), key, value, 0).Err()
	return err
}

// SetWithExpireTime 			插入string类型并设置过期时间
// key 							键
// value 						需要插入的值，内部需要进行序列化
// duration						过期时间
func SetWithExpireTime(key string, value interface{}, duration time.Duration) error {
	err := RedisDB.Set(context.Background(), key, value, duration).Err()
	return err
}

// Get 							获取string类型
// key 							键
// value 						获取的值存储的指针
func Get(key string, value interface{}) error {
	v, err := RedisDB.Get(context.Background(), key).Result()
	val := reflect.ValueOf(v)
	value = val.Interface()
	if err == redis.Nil {
		fmt.Print("key dose not exist")
	} else if err != nil {
		fmt.Printf("get %v failed, err:%v\n", key, err)
		return err
	}
	fmt.Printf("get %v succeed\n, value:%v\n", key, v)
	return nil
}

// GetAndDelete					获取string并删除string类型
// key							键
// value						值
func GetAndDelete(key string, value interface{}) error {
	// Get
	v, err := RedisDB.Get(context.Background(), key).Result()
	if err == redis.Nil {
		fmt.Print("key dose not exist")
	} else if err != nil {
		fmt.Printf("get %v failed, err:%v\n", key, err)
	}
	fmt.Printf("get %v succeed, value:%v\n", key, v)
	val := reflect.ValueOf(v)
	value = val.Interface()
	// Delete
	_, err = RedisDB.Del(context.Background(), key).Result()
	if err == redis.Nil {
		fmt.Print("key dose not exist")
	} else if err != nil {
		fmt.Printf("del %v failed, err:%v\n", key, err)
	}
	fmt.Printf("del %v succeed\n", key)
	return nil
}

// ZSet 						插入set类型
// key						 	键
// value 						需要插入的值
// score 						排序字段名，如果没有该字段则不进行排序
func ZSet(key string, value interface{}, score string) error {
	val := reflect.ValueOf(value)
	addValues := make([]*redis.Z, val.Len())
	fmt.Printf("val.Len = %v\n", val.Len())
	if score != "" {
		scoref64, err := strconv.ParseFloat(score, 64)
		if err != nil {
		}
		for _, value := range val.Interface().([]interface{}) {
			addValues = append(addValues, &redis.Z{
				Score:  scoref64,
				Member: value,
			})
		}
	} else { // 没有设置权值则权值设为默认值0
		for _, value := range val.Interface().([]interface{}) {
			addValues = append(addValues, &redis.Z{
				Score:  0,
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
	val := reflect.ValueOf(score)
	value = val.Interface()
	return err
}

// ZSetWithExpireTime 			插入set类型并设置过期时间
// key							键
// value						需要插入的值
// score						排序字段名，如果没有该字段则不进行排序
// duration						时间
func ZSetWithExpireTime(key string, value interface{}, score string, duration time.Duration) error {
	ok, _ := RedisDB.Expire(context.Background(), key, duration).Result()
	if ok {
		fmt.Println("name 过期时间设置成功", ok)
	} else {
		fmt.Println("name 过期时间设置失败", ok)
	}
	val := reflect.ValueOf(value)
	addValues := make([]*redis.Z, val.Len())
	if score != "" {
		scoref64, err := strconv.ParseFloat(score, 64)
		if err != nil {
		}
		for _, value := range val.Interface().([]interface{}) {
			addValues = append(addValues, &redis.Z{
				Score:  scoref64,
				Member: value,
			})
		}
	} else {
		for _, value := range val.Interface().([]interface{}) {
			addValues = append(addValues, &redis.Z{
				Member: value,
			})
		}
	}
	_, err := RedisDB.ZAdd(context.Background(), key, addValues...).Result()
	return err
}

// Keys 匹配键集
//prefix 键前缀
func Keys(prefix string, value []string) error {
	return nil
}

func GetExpireTime(key string) (time.Duration, error) {
	return 0, nil
}
