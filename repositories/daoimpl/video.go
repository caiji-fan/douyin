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

func (v Video) QueryForUpdate(videoId int) (*po.Video, error) {
	//TODO implement me
	panic("implement me")
}

func (v Video) Begin() (tx *gorm.DB) {
	//TODO implement me
	panic("implement me")
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

func (v Video) QueryBatchIds(videoIds []int) (*[]po.Video, error) {
	//TODO implement me
	panic("implement me")
}

func (v Video) QueryByConditionTimeDESC(condition *po.Video) (*[]po.Video, error) {
	//TODO implement me
	panic("implement me")
}

func (v Video) QueryByLatestTimeDESC(latestTime string) (*[]po.Video, error) {
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
