// Package daoimpl
// @Author shaofan
// @Date 2022/5/13
package daoimpl

import (
	"douyin/entity/po"
	"douyin/repositories"
	"errors"
	"sync"
)

type UserImpl struct {
}

var (
	user     repositories.User
	userOnce sync.Once
)

//单例模式，下次使用如果有这个实例就用这个实例，没有时再创建
func NewUserInstance() repositories.User {
	userOnce.Do(func() {
		user = UserImpl{}
	})
	return user
}
func (UserImpl) QueryById(userId int) (*po.User, error) {
	user := po.User{}
	DB.Where("id = ?", userId).Find(&user)
	if user.Name == "" { //测试如果这个人的名字如果为空的话就是没查到，也可以识别的字段如id
		return &user, errors.New("查询不到该用户！！！")
	}
	return &user, nil
}
func (UserImpl) Insert(user *po.User) error {
	err := DB.Create(user).Error
	return err
}
func (UserImpl) QueryBatchIds(userIds *[]int) (*[]po.User, error) {
	userList := []po.User{}
	DB.Where("id IN ?", *userIds).Find(&userList)
	var err error = nil
	if len(userList) == 0 {
		err = errors.New("查无此人")
	}
	return &userList, err
}
