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
	DB.Where("id = ?", userId).Find(&user) //根据id查数据
	if user.Name == "" {                   //测试如果这个人的名字如果为空的话就是没查到，也可以识别的字段如id
		return &user, errors.New("查询不到该用户！！！")
	}
	return &user, nil
}
func (UserImpl) Insert(user *po.User) (int, error) {
	var ids []int    //插入数据并且返回id
	tx := DB.Begin() //因为同时要插入数据兼返回主键，这里涉及两个数据库操作，所以这里要用手动事务
	err := DB.Create(user).Error
	if err != nil {
		tx.Rollback() //回滚
		return -1, errors.New("注册失败")
	}
	DB.Raw("select LAST_INSERT_ID() as id").Pluck("id", &ids)
	tx.Commit() //提交
	return ids[0], err
}
func (UserImpl) QueryBatchIds(userIds *[]int) (*[]po.User, error) {
	userList := []po.User{} //根据一堆id查多个用户数据
	DB.Where("id IN ?", *userIds).Find(&userList)
	var err error = nil
	if len(userList) == 0 {
		err = errors.New("查无此人")
	}
	return &userList, err
}
func (UserImpl) QueryByName(userName string) (*po.User, error) {
	user := po.User{} //根据名字查用户(用于验证用户名是否唯一，验证密码是否正确)
	DB.Where("name = ?", userName).Find(&user)
	if user.Name != "" {
		return &user, errors.New("用户名重复，注册失败")
	}
	return nil, errors.New("用户名或密码输入错误，请重试")
}
func (UserImpl) UpdateFollowById(id int, t int) error {
	DB.Begin()
	user := po.User{}
	err := DB.Where("id=?", id).Find(&user).Error
	if err != nil || user.Name == "" {
		DB.Rollback()
		return errors.New("该用户不存在，关注数更改失败")
	}
	follow_count := t + user.FollowCount
	if follow_count < 0 {
		DB.Rollback()
		return errors.New("数据出错，关注数不可为负")
	}
	DB.Model(user).Where("id=?", id).Update("follow_count", follow_count)
	DB.Commit()
	return nil
}
func (UserImpl) UpdateFollowerById(id int, t int) error {
	DB.Begin()
	user := po.User{}
	err := DB.Where("id=?", id).Find(&user).Error
	if err != nil || user.Name == "" {
		DB.Rollback()
		return errors.New("该用户不存在，粉丝数更改失败")
	}
	follower_count := t + user.FollowerCount
	if follower_count < 0 {
		DB.Rollback()
		return errors.New("数据出错，粉丝数不可为负")
	}
	DB.Model(user).Where("id=?", id).Update("follower_count", follower_count)
	DB.Commit()
	return nil
}
