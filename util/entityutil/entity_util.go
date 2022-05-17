// Package entityutil
// @Author shaofan
// @Date 2022/5/13
// @DESC 实例转换工具
package entityutil

import (
	"douyin/entity/bo"
	"douyin/entity/po"
)

// GetCommentBOS 	获取评论BO实例集
// src				评论PO集
// dest 			评论bo集
func GetCommentBOS(src *[]po.Comment, dest *[]bo.Comment) error {
	return nil
}

// GetVideoBOS 		获取视频BO实例集
// src				视频PO集
// dest				视频BO集
func GetVideoBOS(src *[]po.Video, dest *[]bo.Video) error {
	return nil
}

// GetUserBOS 		获取用户BO实例集
// src				用户PO集
// dest 			用户BO集
func GetUserBOS(users *[]po.User, dest *[]bo.User) error {
	return nil
}

// GetUserBO 		获取单个用户BO实例
// src				用户PO
// dest				用户BO
func GetUserBO(src *[]po.User, dest *[]bo.User) error {
	return nil
}
