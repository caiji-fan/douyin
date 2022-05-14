// Package service
// @Author shaofan
// @Date 2022/5/13
package service

import (
	"douyin/entity/bo"
	"mime/multipart"
)

// Video				视频业务接口
type Video interface {
	// Feed 			获取feed流
	// userId 			用户id，可以为空
	// isLogin 			用户是否登录，用户未登录时用户id无效
	// @return 			视频列表
	Feed(userId int, isLogin bool) ([]bo.Video, error)

	// Publish 			发布视频
	// file 			视频文件
	// userId 			用户id
	Publish(file *multipart.File, userId int) error

	// VideoList 		查看视频发布列表
	// userId			用户id
	// @return			视频列表
	VideoList(userId int) ([]bo.Video, error)
}
