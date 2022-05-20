///**
// * @Author yg
// * @Date 2022-05-15
// * @Description
// **/
package serviceimpl

import (
	"fmt"
	"testing"
)

//func TestMain(t *testing.M) {
//	config.Init()
//	daoimpl.Init()
//	t.Run()
//}
func TestFavorite_FavoriteList(t *testing.T) {
	boVideos, err := NewFavoriteServiceInstance().FavoriteList(1)
	if err != nil {
		panic(err)
	}
	for _, boVideo := range boVideos {
		fmt.Println(boVideo)
	}
}
