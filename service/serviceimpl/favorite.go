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
	"gorm.io/gorm"
	"sync"
)

type Favorite struct {
}

func (f Favorite) Like(favoriteParam *param.Favorite) error {
	var err error
	tx := daoimpl.NewVideoDaoInstance().Begin()
	userId := favoriteParam.UserID
	videoId := favoriteParam.VideoID
	actionType := favoriteParam.ActionType
	wait := sync.WaitGroup{}
	wait.Add(2)
	if actionType == param.DO_LIKE {
		err = doLike(videoId, userId, &wait, tx)
	} else if actionType == param.CANCEL_LIKE {
		err = cancelLike(videoId, userId, &wait, tx)
	}
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (f Favorite) FavoriteList(userId int) ([]bo.Video, error) {
	var poVideos *[]po.Video
	poVideos, err := daoimpl.NewVideoDaoInstance().QueryVideosByUserId(userId)
	if err != nil {
		return nil, err
	}
	var boVideos = make([]bo.Video, len(*poVideos))
	err = entityutil.GetVideoBOS(poVideos, &boVideos)
	if err != nil {
		return nil, err
	}
	return boVideos, nil
}

// 点赞视频
// videoId 视频id
// userId 用户id
// wait 协程计数
// tx 事务操作
func doLike(videoId, userId int, wait *sync.WaitGroup, tx *gorm.DB) error {
	var err error
	//点赞视频
	go func() {
		defer wait.Done()
		err = favoriteDao.Insert(&po.Favorite{VideoId: videoId, UserId: userId})
		if err != nil {
			return
		}
	}()
	//增加视频点赞数
	go func() {
		defer wait.Done()
		var video *po.Video
		video, err = daoimpl.NewVideoDaoInstance().QueryForUpdate(videoId, tx)
		if err != nil {
			return
		}
		video.FavoriteCount = video.FavoriteCount + 1
		err = daoimpl.NewVideoDaoInstance().UpdateByCondition(video, tx, true)
	}()
	wait.Wait()
	if err != nil {
		return err
	}
	return nil
}

// 取消点赞视频
// videoId 视频id
// userId 用户id
// wait 协程计数
// tx 事务操作
func cancelLike(videoId, userId int, wait *sync.WaitGroup, tx *gorm.DB) error {
	var err error
	//取消点赞视频
	go func() {
		defer wait.Done()
		err = favoriteDao.DeleteByCondition(&po.Favorite{VideoId: videoId, UserId: userId})
		if err != nil {
			return
		}
	}()
	//减少视频点赞数
	go func() {
		defer wait.Done()
		var video *po.Video
		video, err = daoimpl.NewVideoDaoInstance().QueryForUpdate(videoId, tx)
		if err != nil {
			return
		}
		video.FavoriteCount = video.FavoriteCount - 1
		err = daoimpl.NewVideoDaoInstance().UpdateByCondition(video, tx, true)
	}()
	wait.Wait()
	if err != nil {
		return err
	}
	return nil
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
