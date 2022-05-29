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
		fmt.Println("error")
		panic(err)
	}
	for _, boVideo := range boVideos {
		fmt.Println(boVideo)
	}
}
func TestRelation_FansList(t *testing.T) {
	boUsers, err := NewRelationServiceInstance().FansList(1)
	if err != nil {
		panic(err)
	}
	for _, boUser := range *boUsers {
		fmt.Println(boUser)
	}
}
func TestRelation_FollowList(t *testing.T) {
	boUsers, err := NewRelationServiceInstance().FollowList(1)
	if err != nil {
		panic(err)
	}
	for _, boUser := range *boUsers {
		fmt.Println(boUser)
	}
}
