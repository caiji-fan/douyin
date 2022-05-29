// Package controller
// @Author shaofan
// @Date 2022/5/13
package controller

import (
	"douyin/entity/myerr"
	"douyin/entity/param"
	"douyin/entity/response"
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
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(myerr.ArgumentInvalid(webutil.GetValidMsg(err, favorite))))
		return
	}
	err = favoriteService.Like(&favorite)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.SystemError)
		return
	}
	ctx.JSON(http.StatusOK, response.Ok)
}

// FavoriteList 	查看点赞列表
func FavoriteList(ctx *gin.Context) {
	var favoriteListParam param.FavoriteList

	err := ctx.ShouldBindQuery(&favoriteListParam)

	fmt.Println("----------------------------------------", favoriteListParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(myerr.ArgumentInvalid(webutil.GetValidMsg(err, favoriteListParam))))
		return
	}

	videoList, err := favoriteService.FavoriteList(favoriteListParam.UserId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.SystemError)
		return
	}
	ctx.JSON(http.StatusOK, response.FavoriteList{
		Response: response.Ok,
		Data:     videoList,
	})
	return
}
