// Package serviceimpl
// @Author shaofan
// @Date 2022/5/13
package serviceimpl

import (
	"douyin/entity/bo"
	"douyin/entity/param"
	"douyin/entity/po"
	"douyin/repositories/daoimpl"
	"douyin/service"
	"douyin/util/entityutil"
	"sync"
)

type Favorite struct {
}

func (f Favorite) Like(favoriteParam *param.Favorite) error {
	userId := favoriteParam.UserID
	videoId := favoriteParam.VideoID
	actionType := favoriteParam.ActionType
	if actionType == 1 {
		err := favoriteDao.Insert(&po.Favorite{VideoId: videoId, UserId: userId})
		if err != nil {
			return err
		}
		//TODO 异步 更新视频的点赞数
	} else if actionType == 2 {
		err := favoriteDao.DeleteByCondition(&po.Favorite{VideoId: videoId, UserId: userId})
		if err != nil {
			return err
		}
		//TODO 异步 更新视频的点赞数
	}
	return nil
}

func (f Favorite) FavoriteList(userId int) ([]bo.Video, error) {
	videosId, err := favoriteDao.QueryVideoIdsByUserId(userId)
	if err != nil {
		return nil, err
	}
	videoDao := daoimpl.NewVideoDaoInstance()
	var pvideos *[]po.Video
	pvideos, err = videoDao.QueryBatchIds(videosId)
	if err != nil {
		return nil, err
	}
	var bvideos *[]bo.Video
	err = entityutil.GetVideoBOS(pvideos, bvideos)
	if err != nil {
		return nil, err
	}
	return *bvideos, nil
}

func (f Favorite) IsFavorite(videoId int, userId int) (bool, error) {
	flag, err := favoriteDao.QueryByCondition(videoId, userId)
	if err != nil {
		return false, err
	}
	return flag, err
}

var (
	favoriteDao  = daoimpl.NewFavoriteDaoInstance()
	favorite     service.Favorite
	favoriteOnce sync.Once
)

func NewFavoriteServiceInstance() service.Favorite {
	favoriteOnce.Do(func() {
		favorite = Favorite{}
	})
	return favorite
}
