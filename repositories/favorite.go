// Package repositories
// @Author shaofan
// @Date 2022/5/13
package repositories

import (
	"douyin/entity/po"
)

// Favorite 点赞持久层接口
type Favorite interface {
	// Insert 					增加点赞
	// favorite 				一条点赞数据
	Insert(favorite *po.Favorite) error

	// QueryVideoIdsByUserId 	通过用户id查询视频id列表
	// userId 					用户id
	// @return 					视频id集
	QueryVideoIdsByUserId(userId int) ([]int, error)

	// DeleteByCondition		条件删除点赞数据
	// favorite					删除条件
	DeleteByCondition(favorite *po.Favorite) error

	// QueryByCondition			查询用户是否点赞视频
	// videoId					视频id
	// userId					用户id
	// @return 					结果true/false
	QueryByCondition(videoId int, userId int) (bool, error)
}
