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
	err := Set("test1", "succeed!!")
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
	fmt.Printf("val: %v", val)
}

// 正常
func TestRedisUtil_GetAndDelete(t *testing.T) {
	var val interface{}
	err := GetAndDelete("test2", &val)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		panic(err)
	}
	fmt.Printf("val: %v", val)
}

// 正常
func TestRedisUtil_SetWithExpireTime(t *testing.T) {
	err := SetWithExpireTime("test2", "t2 not timeout!", 30000000000)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		panic(err)
	}
}

func TestRedisUtil_ZSet(t *testing.T) {
	var zsetSimple = []string{"milk"}
	err := ZSet("ZZZtest2", zsetSimple, "100")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		panic(err)
	}
}

func TestRedisUtil_ZSetWithExpireTime(t *testing.T) {
	err := ZSetWithExpireTime("ZZZtest2", "milk", "100", 30000000000)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		panic(err)
	}
}

// 正常
func TestRedisUtil_ZGet(t *testing.T) {
	var val interface{}
	err := ZGet("ZZZtest1", &val)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		panic(err)
	}
	fmt.Printf("type: %v, val: %v\n", reflect.TypeOf(val), val)
}
