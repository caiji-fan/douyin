// Package daoimpl
// @Author shaofan
// @Date 2022/5/13
package daoimpl

import (
	"douyin/entity/po"
	"douyin/repositories"
	"gorm.io/gorm"
	"sync"
)

type Video struct {
}

func (v Video) QueryVideosByUserId(userId int) (*[]po.Video, error) {
	var poVideos []po.Video
	err := db.Raw("SELECT v.* FROM dy_video v,dy_favorite f WHERE v.`id`= f.`video_id` AND f.`user_id` = ? ORDER BY f.`create_time` DESC", userId).Scan(&poVideos).Error
	if err != nil {
		return nil, err
	}
	return &poVideos, err
}

func (v Video) QueryForUpdate(videoId int, tx *gorm.DB) (*po.Video, error) {
	var video po.Video
	err := tx.Raw("select id,play_url,cover_url,favorite_count,comment_count,author_id,create_time,update_time from dy_video where id =? FOR UPDATE", videoId).Scan(&video).Error
	return &video, err
}

func (v Video) Begin() *gorm.DB {
	return db.Begin()
}

func (v Video) UpdateByCondition(video *po.Video, tx *gorm.DB, isTx bool) error {
	//TODO implement me
	panic("implement me")
}

func (v Video) QueryById(id int) (*po.Video, error) {
	//TODO implement me
	panic("implement me")
}

func (v Video) Insert(video *po.Video) error {
	//TODO implement me
	panic("implement me")
}

func (v Video) QueryBatchIds(videoIds *[]int, size int) ([]po.Video, error) {
	var videos = make([]po.Video, len(*videoIds))
	return videos, db.Where("id in (?)", *videoIds).Order("create_time DESC").Limit(size).Find(&videos).Error
}

func (v Video) QueryByConditionTimeDESC(condition *po.Video) (*[]po.Video, error) {
	//TODO implement me
	panic("implement me")
}

func (v Video) QueryByLatestTimeDESC(latestTime string, size int) (*[]po.Video, error) {
	//TODO implement me
	panic("implement me")
}

var (
	video     repositories.Video
	videoOnce sync.Once
)

func NewVideoDaoInstance() repositories.Video {
	videoOnce.Do(func() {
		video = Video{}
	})
	return video
}
