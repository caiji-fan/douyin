// Package daoimpl
// @Author shaofan
// @Date 2022/5/13
package daoimpl

import (
	"douyin/entity/po"
	"douyin/repositories"
	"errors"
	"gorm.io/gorm"
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
func (UserImpl) Begin() (*gorm.DB, error) {
	tx := db.Begin()
	if tx == nil {
		return nil, errors.New("事务链接获取失败")
	}
	return tx, nil
}
func (UserImpl) QueryById(userId int) (*po.User, error) {
	db1 := db
	user := po.User{}
	db1.Where("id = ?", userId).Find(&user) //根据id查数据
	if user.Name == "" {                    //测试如果这个人的名字如果为空的话就是没查到，也可以识别的字段如id
		return &user, errors.New("查询不到该数据！！！")
	}
	return &user, nil
}
func (UserImpl) Insert(tx *gorm.DB, isTx bool, user *po.User) (int, error) {
	var client *gorm.DB
	if isTx {
		client = tx
	} else {
		client = db
	}
	var err error
	err = client.Omit("id", "create_time", "update_time").Create(user).Error
	if err != nil {
		return -1, errors.New("验证失败")
	}
	return (*user).ID, err
}
func (UserImpl) QueryBatchIds(userIds *[]int) (*[]po.User, error) {
	db1 := db
	userList := []po.User{} //根据一堆id查多个用户数据
	var err error
	db1.Where("id IN ?", *userIds).Find(&userList)
	if len(userList) == 0 {
		err = errors.New("查不到任何数据")
	}
	return &userList, err
}
func (UserImpl) QueryByCondition(user *po.User) (*[]po.User, error) {
	db1 := db
	var err error
	var users []po.User
	if user.ID != 0 {
		db1 = db1.Where(" id=?", user.ID)
	}
	if user.Name != "" {
		db1 = db1.Where(" name=?", user.Name)
	}
	err = db1.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return &users, nil
}
