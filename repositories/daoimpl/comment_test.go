// Package daoimpl
// @Author shaofan
// @Date 2022/5/15
package daoimpl

import (
	"douyin/entity/po"
	"fmt"
	"log"
	"sync"
	"testing"
)

func TestComment_Insert(t *testing.T) {
	commentDao := NewCommentDaoInstance()
	var comment po.Comment
	comment.VideoId = 1
	comment.SenderId = 1
	comment.Content = "这是评论数据"
	comment.Status = po.NORMAL
	err := commentDao.Insert(&comment)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestComment_QueryByCondition(t *testing.T) {
	commentDao := NewCommentDaoInstance()
	var comment po.Comment
	comment.VideoId = 1
	comment.SenderId = 1
	comment.Status = po.NORMAL
	comments, err := commentDao.QueryByCondition(&comment)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(comments)
}

func TestComment_QueryByCondition2(t *testing.T) {
	wait := sync.WaitGroup{}
	wait.Add(9)
	for i := 1; i < 10; i++ {
		go func() {
			defer wait.Done()
			commentDao := NewCommentDaoInstance()
			var comment po.Comment
			comment.VideoId = 2
			comments, err := commentDao.QueryByCondition(&comment)
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println(comments)
		}()
	}
	wait.Wait()
}

func TestComment_UpdateByCondition(t *testing.T) {
	commentDao := NewCommentDaoInstance()
	var comment po.Comment
	comment.ID = 1
	comment.VideoId = 2
	comment.SenderId = 0
	comment.Status = po.DELETE
	err := commentDao.UpdateByCondition(&comment)
	if err != nil {
		log.Fatalln(err)
	}
}
