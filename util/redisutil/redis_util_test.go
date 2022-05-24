package redisutil

import (
	"douyin/config"
	"fmt"
	"reflect"
	"testing"
)

func TestMain(t *testing.M) {
	config.Init()
	Init()
	fmt.Printf("初始化完成\n")
	t.Run()
}

// 正常
func TestRedisUtil_Set(t *testing.T) {
	err := Set("test2", "congratulation! you are succeed!!")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		panic(err)
	}
}

// 正常
func TestRedisUtil_Get(t *testing.T) {
	var val interface{}
	err := Get("test2", &val)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		panic(err)
	}
	fmt.Printf("val: %v\n", val)
}

// 正常，经过客户端测试会发现确实删掉了，但是缓存的原因这里还可以搜索到
func TestRedisUtil_GetAndDelete(t *testing.T) {
	var val interface{}
	err := GetAndDelete("test4", &val)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		panic(err)
	}
	fmt.Printf("val: %v\n", val)
}

// 正常
func TestRedisUtil_SetWithExpireTime(t *testing.T) {
	err := SetWithExpireTime("test2", "t2 not timeout!", 300000000000)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		panic(err)
	}
}

// 正常
func TestRedisUtil_ZSet(t *testing.T) {
	var zsetSimple = []string{"milk", "coffee", "tea"}
	err := ZSet("ZZZtest2", zsetSimple, "100")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		panic(err)
	}
}

// 正常
func TestRedisUtil_ZSetWithExpireTime(t *testing.T) {
	var zsetSimple = []string{"milk", "bear", "tea"}
	err := ZSetWithExpireTime("ZZZtest3", zsetSimple, "100", 10000000000) // 10s
	if err != nil {
		fmt.Printf("err: %v\n", err)
		panic(err)
	}
}

// 正常
func TestRedisUtil_ZGet(t *testing.T) {
	var val interface{}
	err := ZGet("ZZZtest3", &val)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		panic(err)
	}
	fmt.Printf("type: %v, val: %v\n", reflect.TypeOf(val), val)
}
