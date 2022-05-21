// Package daoimpl
// @Author shaofan
// @Date 2022/5/13
package daoimpl

import (
	"douyin/entity/po"
	"douyin/repositories"
	"sync"

	"gorm.io/gorm"
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
	db1 := db
	var video po.Video
	err := db1.Raw("SELECT * FROM dy_video WHERE id = ? FOR UPDATE", videoId).Scan(&video).Error
	return &video, err
}

func (v Video) Begin() (tx *gorm.DB) {
	return db.Begin()
}

func (v Video) UpdateByCondition(video *po.Video, tx *gorm.DB, isTx bool) error {
	var client *gorm.DB
	if isTx {
		client = tx
	} else {
		client = db
	}
	return client.Model(video).Updates(video).Error
}

func (v Video) QueryById(id int) (*po.Video, error) {
	db1 := db
	var video po.Video
	if id != 0 {
		db1 = db1.Where("id = ?", id)
	}
	err := db1.Find(&video).Error
	return &video, err
}

func (v Video) Insert(video *po.Video) error {
	return db.Omit("ID", "create_time", "update_time").Create(video).Error
}

func (v Video) QueryBatchIds(videoIds []int) (*[]po.Video, error) {
	db1 := db
	videoList := make([]po.Video, len(videoIds))
	err := db1.Where("id IN ?", videoIds).Find(&videoList).Error
	return &videoList, err
}

func (v Video) QueryByConditionTimeDESC(condition *po.Video) (*[]po.Video, error) {
	db1 := db
	var videos []po.Video
	if condition.ID != 0 {
		db1 = db1.Where("id = ?", condition.ID)
	}
	if condition.AuthorId != 0 {
		db1 = db1.Where("author_id = ?", condition.AuthorId)
	}
	if condition.Title != "" {
		db1 = db1.Where("name = ?", condition.Title)
	}
	err := db1.Order("update_time desc").Find(&videos).Error
	return &videos, err
}

func (v Video) QueryByLatestTimeDESC(latestTime string) (*[]po.Video, error) {
	db1 := db
	var videos []po.Video
	if latestTime != "" {
		db1 = db1.Where("update_time < ?", latestTime)
	}
	err := db1.Order("update_time desc").Find(&videos).Error
	return &videos, err
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
