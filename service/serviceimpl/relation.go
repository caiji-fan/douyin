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
	"douyin/util/rabbitutil"
	"sync"
)

type Relation struct {
}

func (r Relation) Follow(relationParam *param.Relation) error {
	userId := relationParam.UserID
	toUserId := relationParam.ToUserID
	actionType := relationParam.ActionType
	if actionType == 1 {
		err := relationDao.Insert(&po.Follow{FollowId: toUserId, FollowerId: userId})
		if err != nil {
			return err
		}
		//使用消息队列 异步 增加user_id用户的关注量  增加toUserId用户的粉丝量
		err1 := rabbitutil.ChangeFollowNum(userId, toUserId, true)
		if err1 != nil {
			return err1
		}
	} else {
		err := relationDao.DeleteByCondition(&po.Follow{FollowId: toUserId, FollowerId: userId})
		if err != nil {
			return err
		}
		//使用消息队列 异步 减少user_id用户的关注量  减少toUserId用户的粉丝量
		err1 := rabbitutil.ChangeFollowNum(userId, toUserId, false)
		if err1 != nil {
			return err1
		}
	}
	return nil
}

func (r Relation) FollowList(userId int) ([]bo.User, error) {
	usersId, err := relationDao.QueryFollowIdByFansId(userId)
	if err != nil {
		return nil, err
	}
	var pusers *[]po.User
	pusers, err = daoimpl.NewUserDaoInstance().QueryBatchIds(&usersId)
	if err != nil {
		return nil, err
	}
	var busers *[]bo.User
	err = entityutil.GetUserBO(pusers, busers)
	if err != nil {
		return nil, err
	}
	return *busers, nil
}

func (r Relation) FansList(userId int) ([]bo.User, error) {
	usersId, err := relationDao.QueryFansIdByFollowId(userId)
	if err != nil {
		return nil, err
	}
	var pusers *[]po.User
	pusers, err = daoimpl.NewUserDaoInstance().QueryBatchIds(&usersId)
	if err != nil {
		return nil, err
	}
	var busers *[]bo.User
	err = entityutil.GetUserBO(pusers, busers)
	if err != nil {
		return nil, err
	}
	return *busers, nil
}

func (r Relation) IsFollow(followId int, followerId int) (bool, error) {
	flag, err := relationDao.QueryByCondition(followId, followerId)
	if err != nil {
		return false, err
	}
	return flag, err
}

var (
	relationDao  = daoimpl.NewRelationDaoInstance()
	relation     service.Relation
	relationOnce sync.Once
)

func NewRelationServiceInstance() service.Relation {
	relationOnce.Do(func() {
		relation = Relation{}
	})
	return relation
}
