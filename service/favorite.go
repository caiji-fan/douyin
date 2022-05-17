// Package service
// @Author shaofan
// @Date 2022/5/13
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

	// IsFavorite				查询用户是否点赞视频
	// videoId					视频id
	// userId					用户id
	// @return 					结果true/false
	IsFavorite(videoId int, userId int) (bool, error)
}
