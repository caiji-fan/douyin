// Package controller
// @Author shaofan
// @Date 2022/5/13
package controller

import (
	"douyin/entity/param"
	"douyin/service/serviceimpl"
	"douyin/util/webutil"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Comment 			评论
func Comment(ctx *gin.Context) {
	var commentParam param.Comment

	err := ctx.ShouldBindQuery(&commentParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status_code": 1,
			"status_msg":  fmt.Sprintf("参数错误 %v", webutil.GetValidMsg(err, commentParam)),
		})
		return
	}
	err = serviceimpl.NewCommentServiceInstance().Comment(&commentParam)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status_code": 1,
			"status_msg":  "系统维护中",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status_code": 0, "status_msg": "ok"})
}

// CommentList 		查看评论列表
func CommentList(ctx *gin.Context) {
	var commentListParam param.CommentList
	err := ctx.ShouldBindQuery(&commentListParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status_code":  1,
			"status_msg":   fmt.Sprintf("参数错误 %v", webutil.GetValidMsg(err, commentListParam)),
			"comment_list": nil,
		})
		return
	}
	commentList, err := serviceimpl.NewCommentServiceInstance().CommentList(commentListParam.VideoId)
	ctx.JSON(http.StatusOK, gin.H{
		"status_code":  0,
		"status_msg":   "ok",
		"comment_list": commentList,
	})
}
