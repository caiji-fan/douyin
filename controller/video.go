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
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Feed 		刷视频
func Feed(ctx *gin.Context) {

}

// Publish 		投稿视频
func Publish(ctx *gin.Context) {
	// 通过线程获取投稿人id
	authorId, err := strconv.Atoi(config.Config.ThreadLocal.Keys.UserId)

	// 通过请求参数获取视频标题
	var videoParm param.Video
	err = ctx.ShouldBindQuery(&videoParm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(myerr.ArgumentInvalid(webutil.GetValidMsg(err, videoParm))))
		return
	}

	// 获取视频本地地址并处理
	video, err := ctx.FormFile("video")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(myerr.VideoNotFound))
		return
	}
	videoPath := filepath.Base(video.Filename)
	videoFinalName := fmt.Sprintf("%d_%s", authorId, videoPath)
	videoSaveFile := filepath.Join("./public/", videoFinalName)
	if err := ctx.SaveUploadedFile(video, videoSaveFile); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(myerr.ArgumentInvalid(webutil.GetValidMsg(err, video))))
		return
	}
	vFile, err := video.Open()

	// 获取封面本地地址并处理
	cover, err := ctx.FormFile("cover")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(myerr.ArgumentInvalid(webutil.GetValidMsg(err, cover))))
		return
	}
	coverPath := filepath.Base(cover.Filename)
	coverFinalName := fmt.Sprintf("%d_%s", authorId, coverPath)
	coverSaveFile := filepath.Join("./public/", coverFinalName)
	if err := ctx.SaveUploadedFile(video, coverSaveFile); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(myerr.ArgumentInvalid(webutil.GetValidMsg(err, cover))))
		return
	}
	cFile, err := cover.Open()

	err = serviceimpl.Publish(&vFile, &cFile, authorId)
	if err != nil {
		ctx.JSON(http.StatusProxyAuthRequired, response.ErrorResponse(err))
	}
	ctx.JSON(http.StatusOK, response.Video{
		Response: response.Ok,
	})
}

// VideoList 	查看视频发布列表
func VideoList(ctx *gin.Context) {

}
