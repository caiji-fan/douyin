package serviceimpl

import (
	"douyin/entity/bo"
	"douyin/util/redisutil"
	"fmt"
	"log"
	"testing"
	"time"
)

func TestVideo_Feed(t *testing.T) {
	redisutil.Init()
	feed, i, err := NewVideoServiceInstance().Feed(1, true, time.Now().UnixMilli())
	if err != nil {
		panic(err)
	}
	fmt.Println(feed)
	fmt.Println(i)
}

func TestVideo_Publish(t *testing.T) {
}

func TestVideo_VideoList(t *testing.T) {

}

// pass
func TestVideo_mergeBox(t *testing.T) {
	var inbox = make([]bo.Feed, 10)
	var outbox = make([]bo.Feed, 10)
	for i := 0; i < 10; i++ {
		inbox[9-i] = bo.Feed{VideoId: 1, CreateTime: time.Now().AddDate(0, 0, i)}
		time.Sleep(100)
		outbox[9-i] = bo.Feed{VideoId: 2, CreateTime: time.Now().AddDate(0, 0, i)}
		time.Sleep(1000)
	}
	box, err := mergeBox(&inbox, &outbox, time.Now().UnixMilli())
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(box)
}

// pass
func TestVideo_mergeFeeds(t *testing.T) {
	var feed1 = make([]bo.Feed, 10)
	var feed2 = make([]bo.Feed, 10)
	for i := 0; i < 10; i++ {
		feed1[9-i] = bo.Feed{VideoId: 1, CreateTime: time.Now().AddDate(0, 0, i)}
		time.Sleep(100)
		feed2[9-i] = bo.Feed{VideoId: 2, CreateTime: time.Now().AddDate(0, 0, i)}
		time.Sleep(1000)
	}
	feed, err := mergeFeeds(&feed1, &feed2)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(feed)
	for _, v := range feed {
		fmt.Println(v.CreateTime.UnixMilli())
	}
	for i := 0; i < len(feed)-2; i++ {

		if feed[i].CreateTime.UnixMilli() < feed[i+1].CreateTime.UnixMilli() {
			log.Fatalln("--------------err-------------------")
		}
	}
}
