// Package repositories
// @Author shaofan
// @Date 2022/5/13
package repositories

import (
	"douyin/entity/po"
)

// Video 视频持久层接口
type Video interface {
	// Insert 						插入
	// video 						视频信息
	Insert(video *po.Video) error

	// QueryBatchIds 				id批量查询
	// videoIds 					视频id集
	// @return 						视频列表
	QueryBatchIds(videoIds []int) (*[]po.Video, error)

	// QueryByConditionTimeDESC 	条件查询并按时间倒序排列
	// condition 					字段匹配查询条件
	// @return 						视频列表
	QueryByConditionTimeDESC(condition *po.Video) (*[]po.Video, error)

	// QueryByLatestTimeDESC 		查询某个时间点之前的视频，时间倒序
	// latestTime					上一次最有一条视频时间
	// @return 						视频列表
	QueryByLatestTimeDESC(latestTime string) (*[]po.Video, error)
}
