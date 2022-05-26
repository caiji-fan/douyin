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
	err := Set("test3", "congratulation! you are succeed!!")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		panic(err)
	}
}

// 正常
func TestRedisUtil_Get(t *testing.T) {
	var val string
	err := Get("test3", &val)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		panic(err)
	}
	fmt.Printf("val: %v\n", val)
}

// 正常，经过客户端测试会发现确实删掉了，但是缓存的原因这里还可以搜索到
func TestRedisUtil_GetAndDelete(t *testing.T) {
	var val interface{}
	err := GetAndDelete("test3", &val)
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

// 弃用
func TestRedisUtil_ZSet(t *testing.T) {
	var zsetSimple = []string{"milk", "coffee", "tea"}
	err := ZSet("ZZZtest3", zsetSimple, "100")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		panic(err)
	}
}

//
func TestRedisUtil_ZSetV2(t *testing.T) {
	var zsetSimple = map[string]float64{"apple": 91, "HuaWei": 90, "xiaomi": 85, "redmi": 88}
	err := ZSetV2("ZZZtest3", zsetSimple)
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
	val := make([]string, 0)
	err := ZGet("ZZZtest3", &val)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		panic(err)
	}
	fmt.Printf("type: %v, val: %v\n", reflect.TypeOf(val), val)
}
