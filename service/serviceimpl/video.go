// Package serviceimpl
// @Author shaofan
// @Date 2022/5/13
package serviceimpl

import (
	"douyin/entity/bo"
	"douyin/entity/po"
	"douyin/entity/response"
	"douyin/repositories/daoimpl"
	"douyin/util/entityutil"
	"douyin/util/obsutil"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/gin-gonic/gin"
)

type Video struct {
}

var wg sync.WaitGroup

// Publish check token then save upload file to public directory
func (v Video) Publish(c *gin.Context, video *multipart.FileHeader, cover *multipart.FileHeader, userId int, title string) {
	// 视频、封面本地保存
	videoname := filepath.Base(video.Filename)
	videoName := fmt.Sprintf("%d_%s", userId, videoname)
	videoSaveFile := filepath.Join("./public/dy/video", videoName)
	if err := c.SaveUploadedFile(video, videoSaveFile); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(err))
	}

	covername := filepath.Base(cover.Filename)
	coverName := fmt.Sprintf("%d_%s", userId, covername)
	coverSaveFile := filepath.Join("./public/dy/cover", coverName)
	if err := c.SaveUploadedFile(video, coverSaveFile); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(err))
	}
	var videoDB daoimpl.Video
	dbinstance := po.Video{
		PlayUrl:       "./public/dy/video" + videoName,
		CoverUrl:      "./public/dy/cover" + coverName,
		FavoriteCount: 0,
		CommentCount:  0,
		AuthorId:      userId,
		Title:         title,
	}
	videoDB.Insert(&dbinstance)

	wg.Add(2)
	// 消息队列异步上传视频， 并更新视频、封面的URL信息， 删除本地视频
	go func() {
		oldVideoUrl := "./public/dy/video" + videoName
		oldCoverUrl := "./public/dy/cover" + coverName
		tx := daoimpl.NewVideoDaoInstance().Begin()
		var videoDB daoimpl.Video
		videourl, err := obsutil.Upload(videoName, "dy-video")
		if err != nil {
			c.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		}
		coverurl, err := obsutil.Upload(coverName, "dy-cover")
		if err != nil {
			c.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		}
		dbinstance := po.Video{
			PlayUrl:       videourl,
			CoverUrl:      coverurl,
			FavoriteCount: 0,
			CommentCount:  0,
			AuthorId:      userId,
			Title:         title,
		}
		videoDB.UpdateByCondition(&dbinstance, tx, true)

		err = os.Remove(oldVideoUrl)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		}
		err = os.Remove(oldCoverUrl)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		}
		wg.Done()
	}()

	// 消息队列异步将视频加入feed流,正确响应
	go func() {

		c.JSON(http.StatusOK, response.PubVideo{
			Response: response.Ok,
		})
		wg.Done()
	}()
}

func (v Video) VideoList(userId int) ([]bo.Video, error) {
	// 查询数据库获取投稿列表
	poVideoList, err := daoimpl.NewVideoDaoInstance().QueryVideosByUserId(userId)
	if err != nil {
		return nil, err
	}
	var boVideoList []bo.Video = make([]bo.Video, len(*poVideoList))
	// po列表转bo
	entityutil.GetVideoBOS(poVideoList, &boVideoList)
	return boVideoList, nil
}
