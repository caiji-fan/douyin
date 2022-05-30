// Package job
// @Author shaofan
// @Date 2022/5/31
package job

import (
	"douyin/config"
	"douyin/util/rabbitutil"
	"douyin/util/redisutil"
	"testing"
)

func TestMain(t *testing.M) {
	config.Init()
	t.Run()
}

// pass
func TestClearOutBox(t *testing.T) {
	redisutil.Init()
	clearOutBox()
}

func TestHandlerErrorMSG(t *testing.T) {
	redisutil.Init()
	rabbitutil.Init()
	handleErrorMSG()
}

// pass
func TestClearLocalVideo(t *testing.T) {
	clearLocalVideo()
}
