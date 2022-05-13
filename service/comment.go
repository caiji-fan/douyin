// Package service
// @Author shaofan
// @Date 2022/5/13
// @DESC
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
	// videoId 			视频id
	CommentList(videoId int) ([]bo.Video, error)
}
