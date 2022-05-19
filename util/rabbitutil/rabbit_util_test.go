// Package rabbitutil
// @Author shaofan
// @Date 2022/5/15
package rabbitutil

import (
	"douyin/config"
	"log"
	"testing"
)

func TestMain(t *testing.M) {
	config.Init()
	Init()
	t.Run()
}
func TestChangeFollowNum(t *testing.T) {
	err := ChangeFollowNum(1, 2, true)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestFeedVideo(t *testing.T) {
	err := FeedVideo(1)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestUploadVideo(t *testing.T) {
	//myerr := UploadVideo()
	//if myerr != nil {
	//	log.Fatalln(myerr)
	//}
}
