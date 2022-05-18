/**
 * @Author yg
 * @Date 2022-05-15
 * @Description
 **/
package daoimpl

import (
	"douyin/config"
	"douyin/entity/po"
	"fmt"
	"testing"
)

func TestMain(t *testing.M) {
	config.Init()
	Init()
	t.Run()
}
func TestRelationDaoImpl_Insert(t *testing.T) {
	relationDao := NewRelationDaoInstance()
	follow := po.Follow{FollowId: 2, FollowerId: 1}
	err := relationDao.Insert(&follow)
	if err != nil {
		fmt.Println(err)
	}
}
func TestRelationDaoImpl_DeleteByCondition(t *testing.T) {
	relationDao := NewRelationDaoInstance()
	follow := po.Follow{FollowId: 1, FollowerId: 0}
	err := relationDao.DeleteByCondition(&follow)
	if err != nil {
		fmt.Println(err)
	}
}
func TestRelationDaoImpl_QueryFollowIdByFansId(t *testing.T) {
	relationDao := NewRelationDaoInstance()
	userIds, err := relationDao.QueryFollowIdByFansId(3)
	if err != nil {
		fmt.Println(err)
	}
	for _, userId := range userIds {
		fmt.Println(userId)
	}
}
func TestRelationDaoImpl_QueryFansIdByFollowId(t *testing.T) {
	relationDao := NewRelationDaoInstance()
	userIds, err := relationDao.QueryFansIdByFollowId(1)
	if err != nil {
		fmt.Println(err)
	}
	for _, userId := range userIds {
		fmt.Println(userId)
	}
}
func TestRelationDaoImpl_QueryByCondition(t *testing.T) {
	relationDao := NewRelationDaoInstance()
	follows, err := relationDao.QueryByCondition(&po.Follow{FollowId: 1, FollowerId: 2})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(follows)
}
func TestFavoriteDaoImpl_Insert(t *testing.T) {
	favoriteDao := NewFavoriteDaoInstance()
	video := po.Favorite{VideoId: 999, UserId: 1}
	err := favoriteDao.Insert(&video)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func TestFavoriteDaoImpl_DeleteByCondition(t *testing.T) {
	favoriteDao := NewFavoriteDaoInstance()
	video := po.Favorite{VideoId: 0, UserId: 1}
	err := favoriteDao.DeleteByCondition(&video)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func TestFavoriteDaoImpl_QueryVideoIdsByUserId(t *testing.T) {
	favoriteDao := NewFavoriteDaoInstance()
	videoIds, err := favoriteDao.QueryVideoIdsByUserId(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(videoIds)
}
func TestFavoriteDaoImpl_QueryByCondition(t *testing.T) {
	favoriteDao := NewFavoriteDaoInstance()
	favorites, err := favoriteDao.QueryByCondition(&po.Favorite{UserId: 1, VideoId: 95})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(favorites)
}
