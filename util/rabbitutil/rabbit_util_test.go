// Package rabbitutil
// @Author shaofan
// @Date 2022/5/15
package rabbitutil

import (
	"douyin/config"
	"douyin/repositories/daoimpl"
	"douyin/util/redisutil"
	"log"
	"testing"
)

func TestMain(t *testing.M) {
	config.Init()
	Init()
	t.Run()
}

// pass
func TestChangeFollowNum(t *testing.T) {
	err := ChangeFollowNum(1, 2, false)
	if err != nil {
		log.Fatalln(err)
	}
}

// pass
func TestFeedVideo(t *testing.T) {
	err := FeedVideo(1)
	if err != nil {
		log.Fatalln(err)
	}
}

// pass
func TestUploadVideo(t *testing.T) {
	err := UploadVideo(1)
	if err != nil {
		log.Fatalln(err)
	}
}

// pass
func TestDoChangeFollowNum(t *testing.T) {
	daoimpl.Init()
	err := doChangeFollowNum(&ChangeFollowNumBody{UserId: 1, ToUserId: 2, IsFollow: true})
	if err != nil {
		log.Fatalln(err)
	}
}

// pass
func TestDoFeedVideo(t *testing.T) {
	daoimpl.Init()
	redisutil.Init()
	err := doFeedVideo(2)
	if err != nil {
		log.Fatalln(err)
	}
}

// pass
func TestDoUploadVideo(t *testing.T) {
	daoimpl.Init()
	redisutil.Init()
	err := doUploadVideo(1)
	if err != nil {
		log.Fatalln(err)
	}
}
