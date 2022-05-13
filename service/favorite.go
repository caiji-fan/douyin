// Package service
// @Author shaofan
// @Date 2022/5/13
// @DESC
package service

import (
	"douyin/entity/bo"
	"douyin/entity/param"
)

// Favorite				点赞业务接口
type Favorite interface {
	// Like 			点赞操作
	// favoriteParam 	点赞参数
	Like(favoriteParam *param.Favorite) error

	// FavoriteList 	点赞列表
	// userId 			用户id
	FavoriteList(userId int) ([]bo.Video, error)
}
