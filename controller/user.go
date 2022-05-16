// Package controller
// @Author shaofan
// @Date 2022/5/13
package controller

import (
	"douyin/entity/param"
	"douyin/service/serviceimpl"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserControllers struct {
}

// Register 		用户注册
func (usc UserControllers) Register(context *gin.Context) {
	var paramuser param.User = param.User{}
	paramuser.UserName = context.Query("username")
	paramuser.Password = context.Query("password")
	userid, token, err := serviceimpl.NewUserService().Register(paramuser)
	if err != nil { //注册失败
		context.JSON(406, gin.H{
			"status_code": "2",
			"status_msg":  err.Error(),
			"user_id":     userid,
			"token":       token,
		})
	} else { //注册成功
		context.JSON(200, gin.H{
			"status_code": "0",
			"status_msg":  "注册成功",
			"user_id":     userid,
			"token":       token,
		})
	}
}

// Login 			用户登录
func (usc UserControllers) Login(context *gin.Context) {
	userName := context.Query("username")
	passWord := context.Query("password")
	userid, token, err := serviceimpl.NewUserService().Login(userName, passWord)
	if err != nil { //登录失败
		context.JSON(407, gin.H{
			"status_code": "2",
			"status_msg":  err.Error(),
			"user_id":     userid,
			"token":       token,
		})
	} else { //登录成功
		context.JSON(200, gin.H{
			"status_code": "0",
			"status_msg":  "注册成功",
			"user_id":     userid,
			"token":       token,
		})
	}
}

// UserInfo 		查看用户信息
func (usc UserControllers) UserInfo(context *gin.Context) {
	userId := context.Query("user_id")
	id, _ := strconv.Atoi(userId)
	user, err := serviceimpl.NewUserService().UserInfo(id)
	if err != nil {
		context.JSON(405, gin.H{
			"status_code": 1,
			"status_msg":  err.Error(),
			"user":        nil,
		})
	} else {
		context.JSON(200, gin.H{
			"status_code": 1,
			"status_msg":  nil,
			"user":        user,
		})
	}
}
