// Package redisutil
// @Author shaofan
// @Date 2022/5/13
// @DESC redis连接初始化与工具
package redisutil

import (
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

// Set 							插入string类型
// key 							键
// value 						需要插入的值，内部需要进行序列化
func Set(key string, value interface{}) error {
	err := RedisDB.Set(key, value, 0).Err()
	return err
}

// SetWithExpireTime 			插入string类型并设置过期时间
// key 							键
// value 						需要插入的值，内部需要进行序列化
// duration						过期时间
func SetWithExpireTime(key string, value interface{}, duration time.Duration) error {
	err := RedisDB.Set(key, value, duration).Err()
	return err
}

// Get 							获取string类型
// key 							键
// value 						获取的值存储的指针
func Get(key string, value interface{}) error {
	v, err := RedisDB.Get(key).Result()
	if err == redis.Nil {
		fmt.Print("key dose not exist\n")
	} else if err != nil {
		fmt.Printf("get %v failed, err:%v\n", key, err)
		return err
	}
	val := reflect.ValueOf(value).Elem()
	val.Set(reflect.ValueOf(v))
	fmt.Printf("get %v succeed, \nvalue:%v\n", key, v)
	return nil
}

// GetAndDelete					获取string并删除string类型
// key							键
// value						值
func GetAndDelete(key string, value interface{}) error {
	// Get
	v, err := RedisDB.Get(key).Result()
	if err == redis.Nil {
		fmt.Print("key dose not exist\n")
	} else if err != nil {
		fmt.Printf("get %v failed, err:%v\n", key, err)
	}
	val := reflect.ValueOf(value).Elem()
	val.Set(reflect.ValueOf(v))
	fmt.Printf("get %v succeed, \nvalue:%v\n", key, v)
	// Delete
	_, err = RedisDB.Del(key).Result()
	if err == redis.Nil {
		fmt.Print("key dose not exist\n")
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
	addValues := make([]redis.Z, val.Len())
	fmt.Printf("val.Len = %v\n", val.Len())
	if score != "" {
		scoref64, err := strconv.ParseFloat(score, 64)
		if err != nil {
		}
		for i := 0; i < val.Len(); i++ {
			addValues = append(addValues, redis.Z{
				Score:  scoref64,
				Member: val.Index(i).Interface(),
			})
		}
	} else { // 没有设置权值则权值设为默认值0
		for i := 0; i < val.Len(); i++ {
			addValues = append(addValues, redis.Z{
				Member: val.Index(i).Interface(),
			})
		}
	}
	_, err := RedisDB.ZAdd(key, addValues...).Result()
	return err
}

func ZSetV2(key string, value map[string]float64) error {
	addValues := make([]redis.Z, len(value))
	for k, v := range value {
		addValues = append(addValues, redis.Z{
			Score:  v,
			Member: reflect.ValueOf(k).Interface(),
		})
	}
	_, err := RedisDB.ZAdd(key, addValues...).Result()
	return err
}

// ZGet 						获取set类型
// key 							键
// value 						获取的值存储的指针
func ZGet(key string, value interface{}) error {
	v, err := RedisDB.ZRange(key, 0, -1).Result()
	fmt.Printf("zget 获得值：%v", v)
	val := reflect.ValueOf(value).Elem()
	val.Set(reflect.ValueOf(v))
	fmt.Printf("get %v succeed, \nvalue:%v\n", key, v)
	return err
}

// ZSetWithExpireTime 			插入set类型并设置过期时间
// key							键
// value						需要插入的值
// score						排序字段名，如果没有该字段则不进行排序
// duration						时间
func ZSetWithExpireTime(key string, value interface{}, score string, duration time.Duration) error {
	val := reflect.ValueOf(value)
	addValues := make([]redis.Z, val.Len())
	if score != "" {
		scoref64, err := strconv.ParseFloat(score, 64)
		if err != nil {
		}
		for i := 0; i < val.Len(); i++ {
			addValues = append(addValues, redis.Z{
				Score:  scoref64,
				Member: val.Index(i).Interface(),
			})
		}
	} else {
		for i := 0; i < val.Len(); i++ {
			addValues = append(addValues, redis.Z{
				Member: val.Index(i).Interface(),
			})
		}
	}
	_, err := RedisDB.ZAdd(key, addValues...).Result()
	ok, _ := RedisDB.Expire(key, duration).Result()
	if ok {
		fmt.Println("name 过期时间设置成功")
	} else {
		fmt.Println("name 过期时间设置失败")
	}
	return err
}
