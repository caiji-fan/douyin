// Package serviceimpl
// @Author shaofan
// @Date 2022/5/13
package serviceimpl

import (
	"douyin/entity/bo"
	"douyin/entity/myerr"
	"douyin/entity/param"
	"douyin/entity/po"
	"douyin/repositories/daoimpl"
	"douyin/service"
	"douyin/util/entityutil"
	"sync"
)

type Comment struct {
}

func (c Comment) Comment(commentParam *param.Comment, userId int) error {
	err := validVideoExistence(commentParam.VideoID)
	if err != nil {
		return err
	}
	if commentParam.ActionType == param.DO_COMMENT {
		return doComment(commentParam, userId)
	} else {
		return deleteComment(commentParam)
	}
}

// 发布评论
func doComment(commentParam *param.Comment, userId int) error {
	commentDao := daoimpl.NewCommentDaoInstance()
	var comment po.Comment
	comment.SenderId = userId
	comment.VideoId = commentParam.VideoID
	comment.Content = commentParam.CommentText
	comment.Status = po.NORMAL
	return commentDao.Insert(&comment)
}

// 删除评论
func deleteComment(commentParam *param.Comment) error {
	commentDao := daoimpl.NewCommentDaoInstance()
	var comment po.Comment
	comment.ID = commentParam.CommentId
	comment.Status = po.DELETE
	return commentDao.UpdateByCondition(&comment)
}

// 校验视频是否存在
func validVideoExistence(videoId int) error {
	videoDao := daoimpl.NewVideoDaoInstance()
	video, err := videoDao.QueryById(videoId)
	if err != nil {
		return err
	}
	// 视频信息不为nil，初始化id为0，为0表示不存在
	if video.ID == 0 {
		return myerr.VideoNotFound
	}
	return nil
}

func (c Comment) CommentList(videoId int) (*[]bo.Comment, error) {
	err2 := validVideoExistence(videoId)
	if err2 != nil {
		return nil, err2
	}
	//查询
	commentDao := daoimpl.NewCommentDaoInstance()
	var comment = new(po.Comment)
	comment.VideoId = videoId
	comment.Status = po.NORMAL
	comments, err := commentDao.QueryByConditionOrderByTime(comment)
	if err != nil {
		return nil, err
	}
	//转换
	var commentBos []bo.Comment
	err = entityutil.GetCommentBOS(comments, &commentBos)
	if err != nil {
		return nil, err
	}
	return &commentBos, nil
}

var (
	comment     service.Comment
	commentOnce sync.Once
)

// NewCommentServiceInstance 获取评论service实例
func NewCommentServiceInstance() service.Comment {
	commentOnce.Do(func() {
		comment = Comment{}
	})
	return comment
}
