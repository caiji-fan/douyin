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
	// ********************得到token与视频信息********************

	// 通过线程获取投稿人id
	authorId, err := strconv.Atoi(config.Config.ThreadLocal.Keys.UserId)
	if err != nil {

	}

	// 通过请求参数获取视频标题
	var videoParm param.Video
	err = ctx.ShouldBindQuery(&videoParm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(myerr.ArgumentInvalid(webutil.GetValidMsg(err, videoParm))))
		return
	}

	// 获取视频本地地址
	video, err := ctx.FormFile("video")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(myerr.VideoNotFound))
		return
	}

	// 获取封面本地地址
	cover, err := ctx.FormFile("cover")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(myerr.ArgumentInvalid(webutil.GetValidMsg(err, cover))))
		return
	}

	var v serviceimpl.Video
	err = v.Publish(video, cover, authorId, videoParm.Title)
	if err != nil {
		ctx.JSON(http.StatusProxyAuthRequired, response.ErrorResponse(err))
	}
	ctx.JSON(http.StatusOK, response.PubVideo{
		Response: response.Ok,
	})
}

// VideoList 	查看视频发布列表
func VideoList(ctx *gin.Context) {
	// 通过请求参数获取对象id
	var videoList param.VideoList
	err := ctx.ShouldBindQuery(&videoList)
	if err != nil {
		return
	}

	var v serviceimpl.Video
	boVideos, err := v.VideoList(videoList.UserID)
	if err != nil {

	}
	ctx.JSON(http.StatusOK, response.VideoList{
		Response:  response.Ok,
		VideoList: boVideos,
	})
}
