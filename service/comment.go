// Package service
// @Author shaofan
// @Date 2022/5/13
package service

import (
	"douyin/entity/bo"
	"douyin/entity/param"
)

// Comment 				评论业务接口
type Comment interface {
	// Comment 			评论操作
	// commentParam 	评论参数
	Comment(commentParam *param.Comment) error

	// CommentList 		查看评论列表
	// commentListParam 视频id参数
	CommentList(commentListParam *param.CommentList) (*[]bo.Comment, error)
}
