// Package controller
// @Author shaofan
// @Date 2022/5/13
package controller

import (
	"douyin/entity/param"
	"douyin/service/serviceimpl"
	"douyin/util/webutil"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var relationService = serviceimpl.NewRelationServiceInstance()

// Follow 		关注与取关
func Follow(ctx *gin.Context) {

	var relation param.Relation

	err := ctx.ShouldBindQuery(&relation)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"status_msg":  fmt.Sprintf("参数错误 %v", webutil.GetValidMsg(err, relation)),
		})
		return
	}

	err = relationService.Follow(&relation)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"status_code": http.StatusForbidden,
			"status_msg":  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "成功",
	})
}

// FollowList 	查看关注列表
func FollowList(ctx *gin.Context) {
	var followListParam param.FollowList
	err := ctx.ShouldBindQuery(&followListParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"status_msg":  fmt.Sprintf("参数错误 %v", webutil.GetValidMsg(err, followListParam)),
			"user_list":   nil,
		})
		return
	}
	userList, err := relationService.FollowList(followListParam.UserID)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"status_code": http.StatusForbidden,
			"status_msg":  err.Error(),
			"user_list":   nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "成功",
		"user_list":   userList,
	})
	return
}

// FansList 	查看粉丝列表
func FansList(ctx *gin.Context) {
	var followListParam param.FollowList
	err := ctx.ShouldBindQuery(&followListParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"status_msg":  fmt.Sprintf("参数错误 %v", webutil.GetValidMsg(err, followListParam)),
			"user_list":   nil,
		})
		return
	}
	userList, err := relationService.FollowList(followListParam.UserID)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"status_code": http.StatusForbidden,
			"status_msg":  err.Error(),
			"user_list":   nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "成功",
		"user_list":   userList,
	})
	return
}
