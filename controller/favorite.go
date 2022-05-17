// Package controller
// @Author shaofan
// @Date 2022/5/13
package controller

import (
	"douyin/entity/param"
	"douyin/service/serviceimpl"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var favoriteService = serviceimpl.NewFavoriteServiceInstance()

// Like 			点赞与取消赞操作
func Like(ctx *gin.Context) {

	var favorite param.Favorite

	err := ctx.ShouldBindQuery(&favorite)

	if err != nil {
		ctx.String(http.StatusBadRequest, "参数错误 %v", GetValidMsg(err, favorite))
		return
	}
	err = favoriteService.Like(&favorite)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"status_msg":  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "成功",
	})
}

// FavoriteList 	查看点赞列表
func FavoriteList(ctx *gin.Context) {
	user_id, _ := strconv.Atoi(ctx.Query("user_id"))
	videoList, err := favoriteService.FavoriteList(user_id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"status_msg":  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "成功",
		"video_list":  videoList,
	})
	return
}
