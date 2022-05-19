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
	"douyin/util/encryptionutil"
	"douyin/util/entityutil"
	"douyin/util/jwtutil"
	"errors"
	"sync"
)

type UserServiceImpl struct {
}

var (
	user     service.User
	userOnce sync.Once
)

func NewUserService() service.User {
	userOnce.Do(func() {
		user = UserServiceImpl{}
	})
	return user
}
func (UserServiceImpl) UserInfo(userId int) (*bo.User, error) {
	userBo := bo.User{}
	userPo, err := daoimpl.NewUserDaoInstance().QueryById(userId) //调用dao层根据id查用户方法
	if err != nil {
		return nil, err
	}
	if userPo == nil {
		return nil, errors.New("查无此人")
	}
	err = entityutil.GetUserBO(userPo, &userBo)
	if err != nil {
		return nil, err
	}
	return &userBo, nil
}
func (UserServiceImpl) Register(userParam param.User) (int, string, error) {
	userPo := po.User{}
	userPo.Name = userParam.UserName
	users, err := daoimpl.NewUserDaoInstance().QueryByCondition(&userPo)
	if err != nil {
		return 0, "", err
	}
	if len(*users) != 0 {
		return 0, "", errors.New("用户名已存在")
	}
	userPo.Password, err = encryptionutil.Encryption(userParam.Password) //调用md5密码加密工具方法
	if err != nil {
		return -6, "", err
	}
	userPo.FollowCount = 0 //初始关注数和粉丝数都应是0
	userPo.FollowerCount = 0
	userId, err := daoimpl.NewUserDaoInstance().Insert(nil, false, &userPo) //执行插入并返回用户id
	jwt, err := jwtutil.CreateJWT(userId)
	if err != nil {
		return 0, "", err
	}
	//TODO redis存token
	return userId, jwt, nil
}
func (UserServiceImpl) Login(userParam param.User) (int, string, error) {
	userBo := po.User{}
	userBo.Name = userParam.UserName
	users, err := daoimpl.NewUserDaoInstance().QueryByCondition(&userBo)
	if err != nil {
		return 0, "", err
	}
	if len(*users) == 0 {
		return 0, "", errors.New("用户不存在")
	}
	tt, err := encryptionutil.EncryptionCompare(userParam.Password, (*users)[0].Password) //调用md5加密对比工具方法
	if err != nil {
		return 0, "", err
	}
	if tt == false {
		return -1, "", err
	}
	//TODO redis存token
	jwt, err := jwtutil.CreateJWT((*users)[0].ID)
	if err != nil {
		return 0, "", err
	}
	return (*users)[0].ID, jwt, nil
}
