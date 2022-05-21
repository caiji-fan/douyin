// Package serviceimpl
// @Author shaofan
// @Date 2022/5/13
package serviceimpl

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

type Video struct {
}

// Publish check token then save upload file to public directory
func Publish(video *multipart.File, cover *multipart.File, userId int) error {
	// 得到token(略)与视频信息
	// 本地视频临时保存视频文件
	// 消息队列异步上传视频
	// 视频信息入库
	// 消息队列异步将视频加入feed流
	// 正确相应

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
