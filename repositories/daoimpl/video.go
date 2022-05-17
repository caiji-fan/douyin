// Package daoimpl
// @Author shaofan
// @Date 2022/5/13
package daoimpl

import (
	"douyin/entity/po"
	"douyin/repositories"
	"sync"
)

type Video struct {
}

func (v Video) UpdateByCondition(video *po.Video) error {
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
