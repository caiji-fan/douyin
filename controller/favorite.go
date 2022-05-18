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

var favoriteService = serviceimpl.NewFavoriteServiceInstance()

// Like 			点赞与取消赞操作
func Like(ctx *gin.Context) {

	var favorite param.Favorite

	err := ctx.ShouldBindQuery(&favorite)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"status_msg":  fmt.Sprintf("参数错误 %v", webutil.GetValidMsg(err, favorite)),
		})
		return
	}
	err = favoriteService.Like(&favorite)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status_code": http.StatusInternalServerError,
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
	var favoriteListParam param.FavoriteList

	err := ctx.ShouldBindQuery(&favoriteListParam)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"status_msg":  fmt.Sprintf("参数错误 %v", webutil.GetValidMsg(err, favoriteListParam)),
		})
		return
	}

	videoList, err := favoriteService.FavoriteList(favoriteListParam.UserId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status_code": http.StatusInternalServerError,
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
