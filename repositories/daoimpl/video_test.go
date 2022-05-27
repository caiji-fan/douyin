package daoimpl

import (
	"douyin/entity/po"
	"fmt"
	"log"
	"testing"
)

func TestVideo_QueryBatchIds(t *testing.T) {
	videoDao := NewVideoDaoInstance()
	var videoIds = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	videos, err := videoDao.QueryBatchIds(&videoIds, 5)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(videos)
}

func TestVideo_QueryForUpdate(t *testing.T) {
	videoDao := NewVideoDaoInstance()
	tx := videoDao.Begin()
	video, err := videoDao.QueryForUpdate(1, tx)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(video)
}

func TestVideoDaoImpl_Insert(t *testing.T) {
	videoDao := NewVideoDaoInstance()
	video := po.Video{
		PlayUrl:       "ticktok.com",
		CoverUrl:      "baidu.com",
		FavoriteCount: 0,
		CommentCount:  0,
		AuthorId:      1,
		Title:         "测试标题：看到就成功",
	}
	err := videoDao.Insert(nil, &video, false)
	if err != nil {
		fmt.Println(err)
	}
}

func TestVideoDaoImpl_QueryById(t *testing.T) {
	videoDao := NewVideoDaoInstance()
	video := &po.Video{
		PlayUrl:       "fack.com",
		CoverUrl:      "fack.com",
		FavoriteCount: 0,
		CommentCount:  0,
		AuthorId:      -1,
		Title:         "测试标题：看到就失败",
	}
	video, err := videoDao.QueryById(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v\n", video)
}

func TestVideoDaoImpl_QueryVideosByUserId(t *testing.T) {
	videoDao := NewVideoDaoInstance()
	video, err := videoDao.QueryVideosByUserId(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v\n", video)
}
