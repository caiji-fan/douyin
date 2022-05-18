// Package entityutil
// @Author shaofan
// @Date 2022/5/13
// @DESC 实例转换工具
package entityutil

import (
	"douyin/config"
	"douyin/entity/bo"
	"douyin/entity/po"
	"douyin/middleware"
	"douyin/repositories/daoimpl"
	"strconv"
)

// GetCommentBOS 	获取评论BO实例集
// src				评论PO集
// dest 			评论bo集
func GetCommentBOS(src *[]po.Comment, dest *[]bo.Comment) error {
	var i = 0
	var ids = make([]int, len(*src), len(*src)*4)
	var cu = make(map[int][]*po.Comment, len(*src))
	//key是po数据库用户id,value是对应的po评论实体切片(一个用户可能评论了多次)
	for _, sr := range *src {
		//用地址可以更省空间，但是也容易出错
		temp := &sr     //temp是sr地址
		temp1 := *temp  //temp1是temp的实体，也就是sr的数据实体
		temp2 := &temp1 //temp2是temp1的地址，
		cu[sr.SenderId] = append(cu[sr.SenderId], temp2)
		ids[i] = sr.SenderId
		i++
	}
	userList, err := daoimpl.NewUserDaoInstance().QueryBatchIds(&ids)
	if err != nil {
		return err
	}
	for _, userPo := range *userList {
		c1s := cu[userPo.ID]
		//将po数据库user转换为bo业务user
		var userBo = bo.User{}
		//通过下面的自定义方法进行
		err := GetUserBO(&userPo, &userBo)
		if err != nil {
			return err
		}
		for _, c1 := range c1s {
			commentBo := bo.Comment{
				ID:         c1.ID,  //bo评论id
				User:       userBo, //bo业务user对象
				Content:    c1.Content,
				CreateDate: c1.CreateTime,
			}
			*dest = append(*dest, commentBo)
		}
	}
	return nil
}

// GetVideoBOS 		获取视频BO实例集
// src				视频PO集
// dest				视频BO集
func GetVideoBOS(src *[]po.Video, dest *[]bo.Video) error {
	var ids = make([]int, len(*src), len(*src)*4)
	var cu = make(map[int][]*po.Video, len(*src))
	//key是作者id,value是视频切片
	//因为一个作者可能有多个作品，所以同上评论一样
	for i, sr := range *src {
		ids[i] = sr.AuthorId
		//用地址可以更省空间，但是也容易出错
		temp := &sr     //temp是sr地址
		temp1 := *temp  //temp1是temp的实体，也就是sr的数据实体
		temp2 := &temp1 //temp2是temp1的地址，
		cu[sr.AuthorId] = append(cu[sr.AuthorId], temp2)
	}
	userList, err := daoimpl.NewUserDaoInstance().QueryBatchIds(&ids)
	if err != nil {
		return err
	}
	for _, userPo := range *userList {
		c1s := cu[userPo.ID]
		//将po数据库user转换为bo业务user
		var userBo = bo.User{}
		//通过下面的自定义方法进行
		err := GetUserBO(&userPo, &userBo)
		if err != nil {
			return err
		}
		for _, c1 := range c1s {
			videoBo := bo.Video{
				ID:            c1.ID,
				Author:        userBo,
				PlayUrl:       c1.PlayUrl,
				CoverUrl:      c1.CoverUrl,
				FavoriteCount: c1.FavoriteCount,
				CommentCount:  c1.CommentCount,
				Title:         c1.Title,
			}
			user_id := middleware.ThreadLocal.Get().(map[string]string)[config.Config.ThreadLocal.Keys.UserId]
			uid, _ := strconv.Atoi(user_id)
			var fav po.Favorite = po.Favorite{
				VideoId: c1.ID,
				UserId:  uid,
			}
			favo, _ := daoimpl.NewFavoriteDaoInstance().QueryByCondition(&fav)
			if favo == nil {
				videoBo.IsFavorite = false
			} else {
				videoBo.IsFavorite = true
			}
			*dest = append(*dest, videoBo)
		}
	}
	return nil
}

// GetUserBOS 		获取用户BO实例集
// src				用户PO集
// dest 			用户BO集
func GetUserBOS(users *[]po.User, dest *[]bo.User) error {
	for i, user := range *users {
		//直接调用单个user转换
		GetUserBO(&user, &((*dest)[i]))
	}
	return nil
}

// GetUserBO 		获取单个用户BO实例
// src				用户PO
// dest				用户BO
func GetUserBO(src *po.User, dest *bo.User) error {
	/*
		bouser.ID = pouser.ID
		bouser.Name = pouser.Name
		bouser.FollowCount = pouser.FollowerCount
		bouser.FollowerCount = pouser.FollowerCount
		bouser.IsFollow = false
	*/
	(*dest).ID = (*src).ID
	(*dest).Name = (*src).Name
	(*dest).FollowCount = (*src).FollowCount
	(*dest).FollowerCount = (*src).FollowerCount
	user_id := middleware.ThreadLocal.Get().(map[string]string)[config.Config.ThreadLocal.Keys.UserId]
	uid, _ := strconv.Atoi(user_id)
	var pof po.Follow = po.Follow{
		FollowId:   uid,
		FollowerId: (*src).ID,
	}
	pofs, _ := daoimpl.NewRelationDaoInstance().QueryByCondition(&pof)
	if pofs == nil {
		(*dest).IsFollow = false
	} else {
		(*dest).IsFollow = true
	}
	return nil
}
