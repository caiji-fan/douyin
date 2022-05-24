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
	var videosMap = make(map[int][]*po.Video, len(*src))
	//key是作者id,value是视频切片
	//因为一个作者可能有多个作品，所以同上评论一样
	for i, sr := range *src {
		ids[i] = sr.AuthorId
		//用地址可以更省空间，但是也容易出错
		temp := &sr     //temp是sr地址
		temp1 := *temp  //temp1是temp的实体，也就是sr的数据实体
		temp2 := &temp1 //temp2是temp1的地址，
		videosMap[sr.AuthorId] = append(videosMap[sr.AuthorId], temp2)
	}
	userList, err := daoimpl.NewUserDaoInstance().QueryBatchIds(&ids) //把这些视频的所有作者查出来
	if err != nil {
		return err
	}
	//userId := middleware.ThreadLocal.Get().(map[string]string)[config.Config.ThreadLocal.Keys.UserId]
	//uid, err := strconv.Atoi(userId) //uid为当前登录用户id
	//if err != nil {
	//	return err
	//}
	//todo单元测试暂改
	uid := 1 //单元测试用
	favoriteVideoId, err := daoimpl.NewFavoriteDaoInstance().QueryVideoIdsByUserId(uid)
	if err != nil {
		return err
	}
	var favoriteVideoIdMap map[int]int = make(map[int]int, len(favoriteVideoId))
	for _, videoId := range favoriteVideoId {
		favoriteVideoIdMap[videoId] = uid
	} //key=视频id(当前登录用户喜欢的所有视频),value=当前登录用户id
	for _, userPo := range *userList {
		videos := videosMap[userPo.ID]
		//将po数据库user转换为bo业务user
		var userBo = bo.User{}
		//通过下面的自定义方法进行
		err := GetUserBO(&userPo, &userBo)
		if err != nil {
			return err
		}
		for _, video := range videos {
			videoBo := bo.Video{
				ID:            video.ID,
				Author:        userBo,
				PlayUrl:       video.PlayUrl,
				CoverUrl:      video.CoverUrl,
				FavoriteCount: video.FavoriteCount,
				CommentCount:  video.CommentCount,
				Title:         video.Title,
			}
			_, boool := favoriteVideoIdMap[video.ID]
			videoBo.IsFavorite = boool
			*dest = append(*dest, videoBo)
		}
	}
	return nil
}

// GetUserBOS 		获取用户BO实例集
// src				用户PO集
// dest 			用户BO集
func GetUserBOS(users *[]po.User, dest *[]bo.User) error {
	//userId := middleware.ThreadLocal.Get().(map[string]string)[config.Config.ThreadLocal.Keys.UserId]
	//uid, err := strconv.Atoi(userId) //查出目前用户的id
	//if err != nil {
	//	return err
	//}
	uid := 1 //测试用
	allfansId, err := daoimpl.NewRelationDaoInstance().QueryFansIdByFollowId(uid)
	//查出目前用户的所有粉丝
	var fansMap map[int]int = make(map[int]int, len(allfansId))
	if err != nil {
		return err
	}
	for _, fan := range allfansId {
		fansMap[fan] = uid //key=粉丝id；value=目前用户id
	}
	for i, poUser := range *users {
		(*dest)[i].ID = poUser.ID
		(*dest)[i].Name = poUser.Name
		(*dest)[i].FollowCount = poUser.FollowerCount
		(*dest)[i].FollowerCount = poUser.FollowerCount
		_, boool := fansMap[poUser.ID]
		(*dest)[i].IsFollow = boool
	}
	return nil
}

// GetUserBO 		获取单个用户BO实例
// src				用户PO
// dest				用户BO
func GetUserBO(src *po.User, dest *bo.User) error {
	//先处理isFollow
	//userId := middleware.ThreadLocal.Get().(map[string]string)[config.Config.ThreadLocal.Keys.UserId]
	//uid, err := strconv.Atoi(userId)
	//if err != nil {
	//	return err
	//}
	//todo单元测试暂改
	uid := 1 //单元测试用
	var poFollow po.Follow = po.Follow{
		FollowId:   uid,
		FollowerId: (*src).ID,
	}
	poFollows, err := daoimpl.NewRelationDaoInstance().QueryByCondition(&poFollow)
	if err != nil {
		return err
	}
	if poFollows == nil {
		(*dest).IsFollow = false
	} else {
		(*dest).IsFollow = true
	}
	//再处理其他简单的处理
	dest.ID = src.ID
	dest.Name = src.Name
	dest.FollowCount = src.FollowCount
	dest.FollowerCount = src.FollowerCount
	return nil
}
