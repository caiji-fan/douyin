// Package daoimpl
// @Author shaofan
// @Date 2022/5/13
package daoimpl

import (
	"douyin/entity/po"
	"douyin/repositories"
	"sync"
)

type Favorite struct {
}

func (f Favorite) Insert(favorite *po.Favorite) error {
	result := db.Select([]string{"video_id", "user_id"}).Create(favorite)
	err := result.Error
	if err != nil {
		return err
	}
	return nil
}

func (f Favorite) QueryVideoIdsByUserId(userId int) ([]int, error) {
	var videoIds []int
	var favorites []po.Favorite
	result := db.Select("video_id").Where("user_id = ?", userId).Order("create_time desc").Find(&favorites)
	err := result.Error
	if err != nil {
		return nil, err
	}
	for _, video := range favorites {
		videoIds = append(videoIds, video.VideoId)
	}
	return videoIds, nil
}

func (f Favorite) DeleteByCondition(favorite *po.Favorite) error {
	videoId := favorite.VideoId
	userId := favorite.UserId
	result := db.Where("video_id = ? and user_id = ?", videoId, userId).Delete(&po.Favorite{})
	err := result.Error
	if err != nil {
		return err
	}
	return nil
}

func (f Favorite) QueryByCondition(videoId int, userId int) (bool, error) {
	result := db.Where("video_id = ? and user_id = ?", videoId, userId).Find(&po.Favorite{})
	err := result.Error
	if err != nil {
		return false, err
	}
	if result.RowsAffected > 0 {
		return true, nil
	} else {
		return false, nil
	}
}

var (
	favorite     repositories.Favorite
	favoriteOnce sync.Once
)

func NewFavoriteDaoInstance() repositories.Favorite {
	favoriteOnce.Do(func() {
		favorite = Favorite{}
	})
	return favorite
}
