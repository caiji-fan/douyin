// Package daoimpl
// @Author shaofan
// @Date 2022/5/13
package daoimpl

import (
	"douyin/entity/po"
	"douyin/repositories"
	"sync"
)

type Relation struct {
}

func (r Relation) Insert(follow *po.Follow) error {
	result := db.Select([]string{"follow_id", "follower_id"}).Create(follow)
	err := result.Error
	if err != nil {
		return err
	}
	return nil
}

func (r Relation) DeleteByCondition(follow *po.Follow) error {
	followId := follow.FollowId
	followerId := follow.FollowerId
	result := db.Where("follow_id = ? and follower_id = ?", followId, followerId).Delete(&po.Follow{})
	err := result.Error
	if err != nil {
		return err
	}
	return nil
}

func (r Relation) QueryByCondition(followId int, followerId int) (bool, error) {
	result := db.Where("follow_id = ? and follower_id = ?", followId, followerId).Find(&po.Follow{})
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

func (r Relation) QueryFollowIdByFansId(fansId int) ([]int, error) {
	var userIds []int
	var follows []po.Follow
	result := db.Select("follow_id").Where("follower_id = ?", fansId).Order("create_time desc").Find(&follows)
	err := result.Error
	if err != nil {
		return nil, err
	}
	for _, follow := range follows {
		userIds = append(userIds, follow.FollowId)
	}
	return userIds, nil
}

func (r Relation) QueryFansIdByFollowId(followId int) ([]int, error) {
	var userIds []int
	var follows []po.Follow
	result := db.Select("follower_id").Where("follow_id = ?", followId).Order("create_time desc").Find(&follows)
	err := result.Error
	if err != nil {
		return nil, err
	}
	for _, follow := range follows {
		userIds = append(userIds, follow.FollowerId)
	}
	return userIds, nil
}

var (
	relation     repositories.Relation
	relationOnce sync.Once
)

func NewRelationDaoInstance() repositories.Relation {
	relationOnce.Do(func() {
		relation = Relation{}
	})
	return relation
}
