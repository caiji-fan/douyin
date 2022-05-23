package daoimpl

import (
	"douyin/entity/po"
	"fmt"
	"testing"
)

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
	err := videoDao.Insert(&video)
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
