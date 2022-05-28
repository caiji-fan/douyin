// Package daoimpl
// @Author shaofan
// @Date 2022/5/22
package daoimpl

import (
	"douyin/entity/po"
	"fmt"
	"log"
	"testing"
)

func TestFeed_InsertBatch(t *testing.T) {
	feedDao := NewFeedDaoInstance()
	var feeds = make([]po.Feed, 10)
	for i := 0; i < 10; i++ {
		feeds[i] = po.Feed{UserId: 1, VideoId: i}
	}
	fmt.Println(feeds)
	err := feedDao.InsertBatch(&feeds, nil, false)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestFeed_QueryByCondition(t *testing.T) {
	feedDao := NewFeedDaoInstance()
	var feed = po.Feed{UserId: 1}
	feeds, err := feedDao.QueryByCondition(&feed)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(feeds)
}

func TestFeed_DeleteByCondition(t *testing.T) {
	feedDao := NewFeedDaoInstance()
	var feed = []po.Feed{{UserId: 1, VideoId: 1}, {UserId: 1, VideoId: 2}}
	err := feedDao.DeleteByCondition(&feed, nil, false)
	if err != nil {
		log.Fatalln(err)
	}
}
