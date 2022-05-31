// Package rabbitutil
// @Author shaofan
// @Date 2022/5/15
package rabbitutil

import (
	"douyin/config"
	"douyin/entity/rabbitentity"
	"douyin/repositories/daoimpl"
	"douyin/util/redisutil"
	"errors"
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
	err := doChangeFollowNum(&rabbitentity.ChangeFollowNumBody{UserId: 1, ToUserId: 2, IsFollow: true})
	if err != nil {
		log.Fatalln(err)
	}
}

// pass
func TestDoFeedVideo(t *testing.T) {
	daoimpl.Init()
	redisutil.Init()
	err := doFeedVideo(11)
	if err != nil {
		panic(err)
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

func TestHandleError(t *testing.T) {
	redisutil.Init()
	var rabbitMSG = rabbitentity.RabbitMSG[int]{Data: 1, Type: rabbitentity.FEED_VIDEO, ResendCount: 0}
	failOnErrorInt(errors.New("测试"), &rabbitMSG)

	var rabbitMSG2 = rabbitentity.RabbitMSG[rabbitentity.ChangeFollowNumBody]{Data: rabbitentity.ChangeFollowNumBody{UserId: 1, ToUserId: 2, IsFollow: true}, ResendCount: 0}
	failOnErrorChangeFollowNumBody(errors.New("测试"), &rabbitMSG2)
}
