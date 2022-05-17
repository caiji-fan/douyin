// Package daoimpl
// @Author shaofan
// @Date 2022/5/13
package daoimpl

import (
	"douyin/entity/po"
	"douyin/repositories"
	"sync"
)

type User struct {
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
