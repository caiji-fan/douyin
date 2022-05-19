// Package serviceimpl
// @Author shaofan
// @Date 2022/5/15
package serviceimpl

import (
	"douyin/config"
	"douyin/entity/param"
	"douyin/repositories/daoimpl"
	"fmt"
	"log"
	"testing"
)

func TestMain(t *testing.M) {
	// 配置文件初始化
	config.Init()
	// 数据库初始化
	daoimpl.Init()
	t.Run()
}

func TestComment_Comment(t *testing.T) {
	commentService := NewCommentServiceInstance()
	var commentParam = param.Comment{ActionType: 1, VideoID: 1, CommentText: "评论", UserId: 1}
	err := commentService.Comment(&commentParam)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestComment_Comment2(t *testing.T) {
	commentService := NewCommentServiceInstance()
	var commentParam = param.Comment{ActionType: 0, CommentId: 1}
	err := commentService.Comment(&commentParam)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestComment_CommentList(t *testing.T) {
	commentService := NewCommentServiceInstance()
	list, err := commentService.CommentList(1)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(list)
}
