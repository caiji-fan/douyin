// Package daoimpl
// @Author shaofan
// @Date 2022/5/13
package daoimpl

import (
	"douyin/entity/po"
	"douyin/repositories"
	"gorm.io/gorm"
	"sync"
)

type UserImpl struct {
}

var (
	user     repositories.User
	userOnce sync.Once
)

// NewUserDaoInstance 单例模式，下次使用如果有这个实例就用这个实例，没有时再创建
func NewUserDaoInstance() repositories.User {
	userOnce.Do(func() {
		user = UserImpl{}
	})
	return user
}
func (UserImpl) Begin() *gorm.DB {
	return db.Begin()
}
func (UserImpl) QueryById(userId int) (*po.User, error) {
	db1 := db
	user := po.User{}
	err := db1.Where("id = ?", userId).Find(&user).Error //根据id查数据
	return &user, err
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
	return (*user).ID, err
}
func (UserImpl) QueryBatchIds(userIds *[]int) (*[]po.User, error) {
	db1 := db
	userList := []po.User{} //根据一堆id查多个用户数据
	var err error
	err = db1.Where("id IN ?", *userIds).Find(&userList).Error
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
	return &users, err
}
func (i UserImpl) QueryForUpdate(userId int) (*po.User, error) {
	db1 := db
	var poUser po.User
	err := db1.Raw("SELECT id,`name`,follow_count,follower_count,`password`,create_time,update_time FROM dy_user WHERE id=? FOR UPDATE", userId).Scan(&poUser).Error
	return &poUser, err
}
func (i UserImpl) UpdateByCondition(user *po.User, tx *gorm.DB, isTx bool) error {
	var client *gorm.DB
	if isTx {
		client = tx
	} else {
		client = db
	}
	return client.Model(user).Updates(user).Error
}
