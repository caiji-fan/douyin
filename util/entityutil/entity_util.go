// Package entityutil
// @Author shaofan
// @Date 2022/5/13
// @DESC 实例转换工具
package entityutil

import (
	"douyin/entity/bo"
	"douyin/entity/po"
	"douyin/repositories/daoimpl"
)

// GetCommentBOS 	获取评论BO实例集
// src				评论PO集
// dest 			评论bo集
func GetCommentBOS(src *[]po.Comment, dest *[]bo.Comment) error {
	var i int = 0
	var ids []int = make([]int, len(*src), len(*src)*4)
	var cu map[int][]*po.Comment = make(map[int][]*po.Comment, len(*src))
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
	userList, err := daoimpl.NewUserInstance().QueryBatchIds(&ids)
	if err != nil {
		return err
	}
	for _, pouser := range *userList {
		c1s := cu[pouser.ID]
		//将po数据库user转换为bo业务user
		var bouser bo.User = bo.User{}
		//通过下面的自定义方法进行
		GetUserBO(&pouser, &bouser)
		for _, c1 := range c1s {
			commentBo := bo.Comment{
				ID:         c1.ID,  //bo评论id
				User:       bouser, //bo业务user对象
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
	var ids []int = make([]int, len(*src), len(*src)*4)
	var cu map[int][]*po.Video = make(map[int][]*po.Video, len(*src))
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
	userList, err := daoimpl.NewUserInstance().QueryBatchIds(&ids)
	if err != nil {
		return err
	}
	for _, pouser := range *userList {
		c1s := cu[pouser.ID]
		//将po数据库user转换为bo业务user
		var bouser bo.User = bo.User{}
		//通过下面的自定义方法进行
		GetUserBO(&pouser, &bouser)
		for _, c1 := range c1s {
			videoBo := bo.Video{
				ID:            c1.ID,
				Author:        bouser,
				PlayUrl:       c1.PlayUrl,
				CoverUrl:      c1.CoverUrl,
				FavoriteCount: c1.FavoriteCount,
				CommentCount:  c1.CommentCount,
				//TODO 查询dy_favorite表才可得出关系，这里先默认true
				IsFavorite: true,
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
	//TODO 处理IsFollow
	//调用repositories中的relation.go中的方法通过pouserID或者bouserID查询是否有关系
	//这里先写固定值
	(*dest).IsFollow = false
	return nil
}
