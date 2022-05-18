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
	userPo, _ := daoimpl.NewUserDaoInstance().QueryById(userId) //调用dao层根据id查用户方法
	if userPo == nil {
		return nil, errors.New("查无此人")
	}
	err := entityutil.GetUserBO(userPo, &userBo)
	if err != nil {
		return nil, err
	}
	return &userBo, nil
}
func (UserServiceImpl) Register(userParam param.User) (int, string, error) {
	userBo := po.User{}
	userBo.Name = userParam.UserName
	users, err := daoimpl.NewUserDaoInstance().QueryByCondition(&userBo)
	if len(*users) != 0 {
		return 0, "", errors.New("用户名已存在")
	}
	userBo.Password, err = encryptionutil.Encryption(userParam.Password) //调用md5密码加密工具方法
	if err != nil {
		return -6, "", err
	}
	userBo.FollowCount = 0 //初始关注数和粉丝数都应是0
	userBo.FollowerCount = 0
	userid, err := daoimpl.NewUserDaoInstance().Insert(nil, false, &userBo) //执行插入并返回用户id
	//TODO token工具类获取token
	jwt, err := jwtutil.CreateJWT(userid)
	if err != nil {
		return 0, "", err
	}
	return userid, jwt, nil
}
func (UserServiceImpl) Login(userParam param.User) (int, string, error) {
	userBo := po.User{}
	userBo.Name = userParam.UserName
	users, err := daoimpl.NewUserDaoInstance().QueryByCondition(&userBo)
	if len(*users) == 0 {
		return 0, "", errors.New("用户不存在")
	}
	tt, err := encryptionutil.EncryptionCompare(userParam.Password, (*users)[0].Password) //调用md5加密对比工具方法
	if tt == false {
		return -1, "", err
	}
	//TODO token工具类获取token
	jwt, err := jwtutil.CreateJWT((*users)[0].ID)
	return (*users)[0].ID, jwt, nil
}
