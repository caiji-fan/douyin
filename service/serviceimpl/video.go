// Package serviceimpl
// @Author shaofan
// @Date 2022/5/13
package serviceimpl

import (
	"douyin/entity/po"
	"douyin/repositories/daoimpl"
	"douyin/util/obsutil"
	"mime/multipart"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type Video struct {
}

// Publish check token then save upload file to public directory
func (v Video) Publish(video *multipart.FileHeader, cover *multipart.FileHeader, userId int, title string) error {

	videoPath := filepath.Base(video.Filename)
	// videoFinalName := fmt.Sprintf("%d_%s", authorId, videoPath)
	// videoSaveFile := filepath.Join("./public/dy/video", videoFinalName)
	// if err := ctx.SaveUploadedFile(video, videoSaveFile); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, response.ErrorResponse(myerr.ArgumentInvalid(webutil.GetValidMsg(err, video))))
	// 	return
	// }
	// vFile, err := video.Open()

	coverPath := filepath.Base(cover.Filename)
	// coverFinalName := fmt.Sprintf("%d_%s", authorId, coverPath)
	// coverSaveFile := filepath.Join("./public/dy/cover", coverFinalName)
	// if err := ctx.SaveUploadedFile(video, coverSaveFile); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, response.ErrorResponse(myerr.ArgumentInvalid(webutil.GetValidMsg(err, cover))))
	// 	return
	// }
	// cFile, err := cover.Open()

	// 消息队列异步上传视频， 并将视频信息写入库
	go func() {
		var videoDB daoimpl.Video
		videourl, err := obsutil.Upload(videoPath, "dy-video")
		if err != nil {

		}
		coverurl, err := obsutil.Upload(coverPath, "dy-cover")
		if err != nil {

		}
		dbinstance := po.Video{
			PlayUrl:       videourl,
			CoverUrl:      coverurl,
			FavoriteCount: 0,
			CommentCount:  0,
			AuthorId:      userId,
			Title:         title,
		}
		videoDB.Insert(&dbinstance)
	}()
	// 消息队列异步将视频加入feed流
	// 正确响应

	return nil
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	// c.JSON(http.StatusOK, VideoList{
	// 	Response: Response{
	// 		StatusCode: 0,
	// 	},
	// 	VideoList: DemoVideos,
	// })
}
