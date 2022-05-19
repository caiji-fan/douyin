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
		err = rabbitutil.ChangeFollowNum(userId, toUserId, true)
		if err != nil {
			return err
		}
	} else {
		err := relationDao.DeleteByCondition(&po.Follow{FollowId: toUserId, FollowerId: userId})
		if err != nil {
			return err
		}
		//使用消息队列 异步 减少user_id用户的关注量  减少toUserId用户的粉丝量
		err = rabbitutil.ChangeFollowNum(userId, toUserId, false)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r Relation) FollowList(userId int) (*[]bo.User, error) {
	//todo 改为userDao QueryFollow
	usersId, err := relationDao.QueryFollowIdByFansId(userId)
	if err != nil {
		return nil, err
	}
	var userPOS *[]po.User
	userPOS, err = daoimpl.NewUserDaoInstance().QueryBatchIds(&usersId)
	if err != nil {
		return nil, err
	}
	var userBOS *[]bo.User
	err = entityutil.GetUserBOS(userPOS, userBOS)
	if err != nil {
		return nil, err
	}
	return userBOS, nil
}

func (r Relation) FansList(userId int) (*[]bo.User, error) {
	//todo 改为userDao QueryFans
	usersId, err := relationDao.QueryFansIdByFollowId(userId)
	if err != nil {
		return nil, err
	}
	var userPOS *[]po.User
	userPOS, err = daoimpl.NewUserDaoInstance().QueryBatchIds(&usersId)
	if err != nil {
		return nil, err
	}
	var userBOS *[]bo.User
	err = entityutil.GetUserBOS(userPOS, userBOS)
	if err != nil {
		return nil, err
	}
	return userBOS, nil
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
