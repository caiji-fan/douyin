// Package daoimpl
// @Author shaofan
// @Date 2022/5/22
package daoimpl

import (
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
