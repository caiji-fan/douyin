// Package redisutil
// @Author shaofan
// @Date 2022/5/13
// @DESC redis连接初始化与工具
package redisutil

import (
	"douyin/entity/myerr"
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"reflect"
	"time"
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
func ZSet(key string, value []redis.Z) error {
	_, err := RedisDB.ZAdd(key, value...).Result()
	return err
}

// ZSetV2 						插入set数据
// key							键
// value						map映射，映射的键为序列化后的值，值为排序值
func ZSetV2(key string, value map[string]float64) error {
	addValues := make([]redis.Z, len(value))
	for k, v := range value {
		addValues = append(addValues, redis.Z{
			Score:  v,
			Member: k,
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
// duration						时间
func ZSetWithExpireTime(key string, value []redis.Z, duration time.Duration, isTx bool, pipeline redis.Pipeliner) error {
	// 如果处于外部事务内，使用外部事务进行处理
	if isTx {
		_, err := pipeline.ZAdd(key, value...).Result()
		if err != nil {
			return nil
		}
		ok, _ := pipeline.Expire(key, duration).Result()
		if ok {
			log.Println("name 过期时间设置成功")
		} else {
			return myerr.RedisExpireError
		}
	} else {
		//如果不处于外部事务，内部使用事务对插入和过期时间两个指令进行事务处理
		pipeline = RedisDB.TxPipeline()
		_, err := pipeline.ZAdd(key, value...).Result()
		if err != nil {
			return err
		}
		ok, _ := pipeline.Expire(key, duration).Result()
		if ok {
			log.Println("name 过期时间设置成功")
		} else {
			err1 := pipeline.Discard()
			if err1 != nil {
				return err1
			}
			return myerr.RedisExpireError
		}
		if err != nil {
			err1 := pipeline.Discard()
			if err1 != nil {
				return err1
			}
			return err
		}
		_, err = pipeline.Exec()
		if err != nil {
			return err
		}
	}

	return nil
}

// Keys 获取匹配的键集
func Keys(prefix string, keys *[]string) error {
	res := RedisDB.Keys(prefix + "*")
	*keys = res.Val()
	return res.Err()
}

// GetExpireTime 获取的键值的过期时间
func GetExpireTime(key string) (time.Duration, error) {
	res := RedisDB.TTL(key)
	return res.Val(), res.Err()
}

// Lock 加锁
// todo wangyingsong
func Lock(key string, expireTime time.Duration) (bool, error) {
	return false, nil
}

// Unlock 解锁
//todo wangyingsong
func Unlock(key string) error {
	return nil
}

// Begin 开启事务
func Begin() redis.Pipeliner {
	return RedisDB.TxPipeline()
}
