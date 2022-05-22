// Package controller
// @Author shaofan
// @Date 2022/5/13
package controller

import (
	"douyin/config"
	"douyin/entity/myerr"
	"douyin/entity/param"
	"douyin/entity/response"
	"douyin/service/serviceimpl"
	"douyin/util/webutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Feed 		刷视频
func Feed(ctx *gin.Context) {

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

}
