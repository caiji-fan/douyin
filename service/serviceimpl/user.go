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
	"sync"
	"time"
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
func (UserServiceImpl) UserInfo(userId int) (bo.User, error) {
	bouser := bo.User{}
	pouser, err := daoimpl.NewUserInstance().QueryById(userId) //调用dao层根据id查用户方法
	if err != nil {
		return bouser, err
	}
	entityutil.GetUserBO(pouser, &bouser) //调用自定义实体类转换工具
	return bouser, err
}
func (UserServiceImpl) Register(userParam param.User) (int, string, error) {
	var username string = userParam.UserName
	user, err := daoimpl.NewUserInstance().QueryByName(username) //调用dao层根据名字查用户方法
	if user != nil {
		return -101, "", err
	}
	pouser := po.User{}
	pouser.Name = username
	pouser.Password, err = encryptionutil.Encryption(userParam.Password) //调用md5密码加密工具方法
	if err != nil {
		return -6, "", err
	}
	pouser.CreateTime = time.Now().Format("2006-01-02 15:04:05") //插入此格式的现在时间
	pouser.UpdateTime = time.Now().Format("2006-01-02 15:04:05")
	pouser.FollowCount = 0 //初始关注数和粉丝数都应是0
	pouser.FollowerCount = 0
	userid, err := daoimpl.NewUserInstance().Insert(&pouser) //执行插入并返回用户id
	//TODO token工具类获取token
	return userid, "token", err
}
func (UserServiceImpl) Login(userName string, password string) (int, string, error) {
	user, err := daoimpl.NewUserInstance().QueryByName(userName) //调用根据name查用户
	if user == nil {                                             //通过用户名查不到信息，用户名出错
		return -1, "", err
	} //下面直接调用加密密码对比方法
	tt, err := encryptionutil.EncryptionCompare(password, (*user).Password) //调用md5加密对比工具方法
	if tt == false {
		return -1, "", err
	}
	//TODO token工具类获取token
	return (*user).ID, "token", nil
}
