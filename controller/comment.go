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

// Comment 			评论
func Comment(ctx *gin.Context) {
	var commentParam param.Comment
	err := ctx.ShouldBindQuery(&commentParam)
	if err != nil {
		ctx.String(http.StatusBadRequest, "参数错误 %v", GetValidMsg(err, commentParam))
		return
	}
	err = serviceimpl.NewCommentServiceInstance().Comment(&commentParam)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "系统维护中")
	}
	ctx.JSON(http.StatusOK, gin.H{"status_code": 0, "status_msg": "ok"})
}

// CommentList 		查看评论列表
func CommentList(ctx *gin.Context) {
	videoId, err := strconv.Atoi(ctx.Query("video_id"))
	if err != nil {
		ctx.String(http.StatusBadRequest, "参数错误")
	}
	commentList, err := serviceimpl.NewCommentServiceInstance().CommentList(videoId)
	ctx.JSON(http.StatusOK, commentList)
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
