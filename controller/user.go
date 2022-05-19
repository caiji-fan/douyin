// Package controller
// @Author shaofan
// @Date 2022/5/13
package controller

import (
	"douyin/entity/param"
	"douyin/service/serviceimpl"
	"douyin/util/webutil"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Register 		用户注册
func Register(context *gin.Context) {
	var user param.User
	err := context.ShouldBindQuery(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status_code": -1,
			"status_msg":  webutil.GetValidMsg(err, user),
			"user_id":     0,
			"token":       "",
		})
		return
	}
	userId, token, err := serviceimpl.NewUserService().Register(user)
	if err != nil { //注册失败
		context.JSON(http.StatusInternalServerError, gin.H{
			"status_code": "2",
			"status_msg":  err.Error(),
			"user_id":     userId,
			"token":       token,
		})
	} else { //注册成功
		context.JSON(http.StatusOK, gin.H{
			"status_code": 0,
			"status_msg":  "注册成功",
			"user_id":     userId,
			"token":       token,
		})
	}
}

// Login 			用户登录
func Login(context *gin.Context) {
	var user param.User
	err := context.ShouldBindQuery(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status_code": -1,
			"status_msg":  webutil.GetValidMsg(err, user),
			"user_id":     0,
			"token":       "",
		})
		return
	}
	userId, token, err := serviceimpl.NewUserService().Login(user)
	if err != nil { //登录失败
		context.JSON(407, gin.H{
			"status_code": 2,
			"status_msg":  err.Error(),
			"user_id":     userId,
			"token":       token,
		})
	} else { //登录成功
		context.JSON(200, gin.H{
			"status_code": 0,
			"status_msg":  "登录成功",
			"user_id":     userId,
			"token":       token,
		})
	}
}

// UserInfo 		查看用户信息
func UserInfo(context *gin.Context) {
	var userInfoParam param.UserInfo
	err := context.ShouldBindQuery(&userInfoParam)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status_code": -1,
			"status_msg":  webutil.GetValidMsg(err, userInfoParam),
			"user":        nil,
		})
		return
	}
	user, err := serviceimpl.NewUserService().UserInfo(userInfoParam.UserId)
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
