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

type User struct {
}

func (u User) UpdateByCondition(user *po.User, tx *gorm.DB, isTx bool) error {
	//TODO implement me
	panic("implement me")
}

func (u User) Begin() (tx *gorm.DB) {
	//TODO implement me
	panic("implement me")
}

func (u User) QueryForUpdate(userId int) (*po.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u User) Insert(user *po.User) error {
	//TODO implement me
	panic("implement me")
}

func (u User) QueryById(userId int) (*po.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u User) QueryBatchIds(userIds *[]int) (*[]po.User, error) {
	//TODO implement me
	panic("implement me")
}

var (
	user     repositories.User
	userOnce sync.Once
)

func NewUserDaoInstance() repositories.User {
	userOnce.Do(func() {
		user = User{}
	})
	return user
}
