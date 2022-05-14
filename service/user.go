// Package service
// @Author shaofan
// @Date 2022/5/13
package service

import (
	"douyin/entity/bo"
	"douyin/entity/param"
)

// User 			用户业务接口
type User interface {
	// Register 	注册
	// userParam 	用户信息
	// @return 		鉴权token
	Register(userParam param.User) (string, error)

	// Login 		登录
	// userName 	用户名
	// password 	用户密码
	// @return 		鉴权token
	Login(userName string, password string) (string, error)

	// UserInfo 	查看用户信息
	// userId 		用户id
	// @return 		用户信息
	UserInfo(userId int) (bo.User, error)
}
