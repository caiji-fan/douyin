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

	var err error
	tx := daoimpl.NewVideoDaoInstance().Begin()
	userId := favoriteParam.UserID
	videoId := favoriteParam.VideoID
	actionType := favoriteParam.ActionType
	if actionType == 1 {
		wait := sync.WaitGroup{}
		wait.Add(2)
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
			video, err = daoimpl.NewVideoDaoInstance().QueryForUpdate(videoId)
			if err != nil {
				return
			}
			video.FavoriteCount = video.FavoriteCount + 1
			err = daoimpl.NewVideoDaoInstance().UpdateByCondition(video, tx, true)
		}()
		wait.Wait()
		if err != nil {
			tx.Rollback()
			return err
		}
		tx.Commit()
	} else if actionType == 2 {
		wait := sync.WaitGroup{}
		wait.Add(2)
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
			video, err = daoimpl.NewVideoDaoInstance().QueryForUpdate(videoId)
			if err != nil {
				return
			}
			video.FavoriteCount = video.FavoriteCount - 1
			err = daoimpl.NewVideoDaoInstance().UpdateByCondition(video, tx, true)
		}()
		wait.Wait()
		if err != nil {
			tx.Rollback()
			return err
		}
		tx.Commit()
	}
	return nil
}

func (f Favorite) FavoriteList(userId int) ([]bo.Video, error) {
	//todo 调用视频的联查
	videosId, err := favoriteDao.QueryVideoIdsByUserId(userId)
	if err != nil {
		return nil, err
	}
	videoDao := daoimpl.NewVideoDaoInstance()
	var poVideos *[]po.Video
	poVideos, err = videoDao.QueryBatchIds(videosId)
	if err != nil {
		return nil, err
	}
	var boVideos *[]bo.Video
	err = entityutil.GetVideoBOS(poVideos, boVideos)
	if err != nil {
		return nil, err
	}
	return *boVideos, nil
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
