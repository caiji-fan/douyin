// Package controller
// @Author shaofan
// @Date 2022/5/13
package controller

import (
	"douyin/entity/param"
	"douyin/service/serviceimpl"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"reflect"
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

// GetValidMsg  通过错误获取自定义的提示信息
func GetValidMsg(err error, obj interface{}) string {
	//通过反射获取结构体
	getObj := reflect.TypeOf(obj)
	//取得错误信息
	if errs, ok := err.(validator.ValidationErrors); ok {
		//遍历所有校验错误
		for _, e := range errs {
			//遍历结构体中的字段
			for i := 0; i < getObj.NumField(); i++ {
				//当结构体中某个字段和出错的字段相同时，返回字段标签中的msg，这个msg就是自定义的错误提示
				if getObj.Field(i).Name == e.Field() {
					return getObj.Field(i).Tag.Get("msg")
				}
			}
		}
	}
	//如果没有找到该字段直接返回错误
	return err.Error()
}
