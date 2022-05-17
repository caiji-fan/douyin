// Package serviceimpl
// @Author shaofan
// @Date 2022/5/13
package serviceimpl

import (
	"douyin/entity/bo"
	"douyin/entity/param"
	"douyin/entity/po"
	"douyin/repositories/daoimpl"
	"douyin/service"
	"douyin/util/entityutil"
	"errors"
	"sync"
)

type Comment struct {
}

func (c Comment) Comment(commentParam *param.Comment) error {
	//todo 校验视频是否存在
	//err := validVideoExistence(commentParam.VideoID)
	//if err != nil {
	//	return err
	//}
	if commentParam.ActionType == param.DO_COMMENT {
		return doComment(commentParam)
	} else {
		return deleteComment(commentParam)
	}
}

// 发布评论
func doComment(commentParam *param.Comment) error {
	commentDao := daoimpl.NewCommentDaoInstance()
	var comment po.Comment
	comment.SenderId = commentParam.UserId
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
	if video != nil {
		return errors.New("视频不存在")
	}
	return nil
}

func (c Comment) CommentList(videoId int) (*[]bo.Comment, error) {
	//todo 校验视频是否存在
	// 校验
	//err := validVideoExistence(videoId)
	//if err != nil {
	//	return nil, err
	//}
	// 查询
	commentDao := daoimpl.NewCommentDaoInstance()
	var comment = new(po.Comment)
	comment.VideoId = videoId
	comments, err := commentDao.QueryByCondition(comment)
	if err != nil {
		return nil, err
	}
	//转换
	var commentBos *[]bo.Comment
	err = entityutil.GetCommentBOS(comments, commentBos)
	if err != nil {
		return nil, err
	}
	return commentBos, nil
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
