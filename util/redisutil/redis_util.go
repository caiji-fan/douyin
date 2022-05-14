// Package redisutil
// @Author shaofan
// @Date 2022/5/13
// @DESC redis连接初始化与工具
package redisutil

import "time"

// Set 							插入string类型
// key 							键
// value 						需要插入的值，内部需要进行序列化
func Set(key string, value *any) error {
	return nil
}

// SetWithExpireTime 			插入string类型并设置过期时间
// key 							键
// value 						需要插入的值，内部需要进行序列化
// duration						过期时间
func SetWithExpireTime(key string, value *any, duration time.Duration) error {
	return nil
}

// Get 							获取string类型
// key 							键
// value 						获取的值存储的指针
func Get(key string, value *any) error {
	return nil
}

// ZSet 						插入set类型
// key						 	键
// value 						需要插入的值
// score 						排序字段名，如果没有该字段则不进行排序
func ZSet(key string, value *[]any, score string) error {
	return nil
}

// ZGet 						获取set类型
// key 							键
// value 						获取的值存储的指针
func ZGet(key string, value *[]any) error {
	return nil
}
