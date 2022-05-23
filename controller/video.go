// Package controller
// @Author shaofan
// @Date 2022/5/13
package controller

import (
	"douyin/config"
	"douyin/entity/response"
	"douyin/middleware"
	"douyin/service/serviceimpl"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// Feed 		刷视频
func Feed(ctx *gin.Context) {
	var latestTime int64
	var timeParam = ctx.Query("latest_time")
	var err error
	videoService := serviceimpl.NewVideoServiceInstance()
	if timeParam == "" {
		latestTime = time.Now().UnixMilli()
	} else {
		latestTime, err = strconv.ParseInt(timeParam, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, response.SystemError)
			return
		}
	}
	var userId = 0
	var isLogin = false
	if middleware.ThreadLocal.Get() != nil {
		isLogin = true
		userId, err = strconv.Atoi(middleware.ThreadLocal.Get().(map[string]string)[config.Config.ThreadLocal.Keys.UserId])
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, response.SystemError)
			return
		}
	}
	feed, nextTime, err := videoService.Feed(userId, isLogin, latestTime)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, response.Feed{
		VideoList: feed,
		NextTime:  nextTime,
	})
}

// Publish 		投稿视频
func Publish(ctx *gin.Context) {
}

// VideoList 	查看视频发布列表
func VideoList(ctx *gin.Context) {

}
