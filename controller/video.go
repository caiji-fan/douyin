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
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// Feed 		刷视频
func Feed(ctx *gin.Context) {

}

// Publish 		投稿视频
func Publish(ctx *gin.Context) {
	var video param.Video
	err := ctx.ShouldBindQuery(&video)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(myerr.ArgumentInvalid(webutil.GetValidMsg(err, video))))
		return
	}

	data, err := ctx.FormFile("data")
	if err != nil {
		response.ErrorResponse(myerr.VideoNotFound)
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(myerr.ArgumentInvalid(webutil.GetValidMsg(err, data))))
		return
	}
	filename := filepath.Base(data.Filename)

	finalName := fmt.Sprintf("%d_%s", video.AuthorId, filename)
	saveFile := filepath.Join("./public/", finalName)
	if err := ctx.SaveUploadedFile(data, saveFile); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(myerr.ArgumentInvalid(webutil.GetValidMsg(err, data))))
		return
	}

	file, err := data.Open()
	err = serviceimpl.Publish(&file, video.AuthorId)
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
