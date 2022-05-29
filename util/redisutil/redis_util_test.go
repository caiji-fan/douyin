package redisutil

import (
	"douyin/config"
	"douyin/entity/bo"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"reflect"
	"testing"
	"time"
)

func TestMain(t *testing.M) {
	config.Init()
	Init()
	t.Run()
}

type TestRedis struct {
	Value string `json:"value"`
}

type TestRedis2 struct {
	Value string    `json:"value"`
	Date  time.Time `json:"date"`
}

func (t TestRedis) MarshalBinary() ([]byte, error) {
	return json.Marshal(t)
}
func (t TestRedis2) MarshalBinary() ([]byte, error) {
	return json.Marshal(t)
}

// pass
func TestSet(t *testing.T) {
	err := Set("test", TestRedis{Value: "456"})
	if err != nil {
		log.Fatalln(err)
	}
}

// pass
func TestGet(t *testing.T) {
	var val TestRedis
	err := Get[TestRedis]("test", &val)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("val: %v\n", val)
}

// pass
func TestGetAndDelete(t *testing.T) {
	var val TestRedis
	err := GetAndDelete[TestRedis]("test", &val)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("val: %v\n", val)
}

// 正常
func TestSetWithExpireTime(t *testing.T) {
	duration, err := time.ParseDuration("1m")
	if err != nil {
		log.Fatalln(err)
	}
	err = SetWithExpireTime("test", TestRedis{Value: "456"}, duration)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		panic(err)
	}
}

func TestTTL(t *testing.T) {
	ttl, err := TTL("test")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(ttl)
}

// pass
func TestZAdd(t *testing.T) {
	var value = make([]redis.Z, 2)
	time1, _ := time.Parse(config.Config.StandardTime, "2002-2-2")
	time2, _ := time.Parse(config.Config.StandardTime, "2002-2-3")
	value[0] = redis.Z{Score: 1, Member: TestRedis2{Value: "123", Date: time1}}
	value[1] = redis.Z{Score: 1, Member: TestRedis2{Value: "456", Date: time2}}
	err := ZAdd("test", value, false, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

// pass
func TestZRem(t *testing.T) {
	time1, _ := time.Parse(config.Config.StandardTime, "2002-2-2")
	time2, _ := time.Parse(config.Config.StandardTime, "2002-2-3")
	var value = []TestRedis2{{Value: "123", Date: time1}, {Value: "456", Date: time2}}
	err := ZRem[TestRedis2]("test", &value, false, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

// pass
func TestZAddWithExpireTime(t *testing.T) {
	var value = make([]redis.Z, 1)
	value[0] = redis.Z{Score: 1, Member: bo.Feed{VideoId: 1, CreateTime: time.Now()}}
	expireTime, err := time.ParseDuration(config.Config.Redis.ExpireTime.Inbox)
	if err != nil {
		log.Fatalln(err)
	}
	err = ZAddWithExpireTime(config.Config.Redis.Key.Inbox+"1", value, expireTime, false, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

// pass
func TestZRevRange(t *testing.T) {
	val := make([]bo.Feed, 0)
	err := ZRevRange[bo.Feed]("test", &val)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("type: %v, val: %v\n", reflect.TypeOf(val), val)
}

// pass
func TestKeys(t *testing.T) {
	var keys []string
	err := Keys("name", &keys)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(keys)
}

//pass
func TestBegin(t *testing.T) {
	p := Begin()
	fmt.Println(p)
}

func TestLock(t *testing.T) {

}

func TestUnlock(t *testing.T) {

}
