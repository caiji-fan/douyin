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
	err := db.Select([]string{"video_id", "user_id"}).Create(favorite).Error
	if err != nil {
		return err
	}
	return nil
}

func (f Favorite) QueryVideoIdsByUserId(userId int) ([]int, error) {
	videoIds := []int{}
	favorites := []po.Favorite{}
	err := db.Select("video_id").Where("user_id = ?", userId).Find(&favorites).Error
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
	db1 := db
	if videoId != 0 {
		db1 = db1.Where("video_id = ?", videoId)
	}
	if userId != 0 {
		db1 = db1.Where("user_id = ?", userId)
	}
	err := db1.Delete(&po.Favorite{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (f Favorite) QueryByCondition(favorite *po.Favorite) (*[]po.Favorite, error) {
	videoId := favorite.VideoId
	userId := favorite.UserId
	db1 := db
	if videoId != 0 {
		db1 = db1.Where("video_id = ?", videoId)
	}
	if userId != 0 {
		db1 = db1.Where("user_id = ?", userId)
	}
	var favorites []po.Favorite
	err := db1.Find(&favorites).Error
	if err != nil {
		return nil, err
	}
	return &favorites, nil
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
