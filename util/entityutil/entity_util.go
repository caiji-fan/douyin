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
	if *dest == nil || len(*dest) < len(*src) {
		*dest = make([]bo.Comment, len(*src))
	}
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
	var destMap = make(map[int]bo.Comment, len(*src)) //key:评论id;value:bo评论实体
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
				CreateDate: c1.CreateTime.Format("01-02"),
			}
			destMap[commentBo.ID] = commentBo
		}
	}
	for i, comment := range *src {
		(*dest)[i] = destMap[comment.ID]
	}
	return nil
}

// GetVideoBOS 		获取视频BO实例集
// src				视频PO集
// dest				视频BO集
func GetVideoBOS(src *[]po.Video, dest *[]bo.Video) error {
	if *dest == nil || len(*dest) < len(*src) {
		*dest = make([]bo.Video, len(*src))
	}
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
	var isLogin bool = true
	var favoriteVideoIdMap map[int]int
	//todo暂改
	if middleware.ThreadLocal.Get() == nil { //未登录状态
		isLogin = false
	} else { //已登陆状态
		userId := middleware.ThreadLocal.Get().(map[string]string)[config.Config.ThreadLocal.Keys.UserId]
		uid, err := strconv.Atoi(userId) //uid为当前登录用户id
		if err != nil {
			return err
		}
		//todo单元测试暂改
		//uid := 1 //单元测试用
		favoriteVideoId, err := daoimpl.NewFavoriteDaoInstance().QueryVideoIdsByUserId(uid)
		if err != nil {
			return err
		}
		favoriteVideoIdMap = make(map[int]int, len(favoriteVideoId))
		for _, videoId := range favoriteVideoId {
			favoriteVideoIdMap[videoId] = uid
		} //key=视频id(当前登录用户喜欢的所有视频),value=当前登录用户id
	}
	var destMap = make(map[int]bo.Video, len(*src)) //key:视频id;value:bo视频对象
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
			if isLogin {
				_, boool := favoriteVideoIdMap[video.ID]
				videoBo.IsFavorite = boool
			} else {
				videoBo.IsFavorite = false
			}
			destMap[videoBo.ID] = videoBo
		}
	}
	for i, video := range *src {
		(*dest)[i] = destMap[video.ID]
	}
	return nil
}

// GetUserBOS 		获取用户BO实例集
// src				用户PO集
// dest 			用户BO集
func GetUserBOS(users *[]po.User, dest *[]bo.User) error {
	if *dest == nil || len(*dest) < len(*users) {
		*dest = make([]bo.User, len(*users))
	}
	userId := middleware.ThreadLocal.Get().(map[string]string)[config.Config.ThreadLocal.Keys.UserId]
	uid, err := strconv.Atoi(userId) //查出目前用户的id
	if err != nil {
		return err
	}
	//todo测试用
	//uid := 1 //测试用
	allFollowsId, err := daoimpl.NewRelationDaoInstance().QueryFollowIdByFansId(uid)
	//查出目前用户关注的所有的人的id
	var followsMap = make(map[int]int, len(allFollowsId))
	if err != nil {
		return err
	}
	for _, follow := range allFollowsId {
		followsMap[follow] = uid //key=关注的人的id；value=目前用户id
	}
	for i, poUser := range *users {
		(*dest)[i].ID = poUser.ID
		(*dest)[i].Name = poUser.Name
		(*dest)[i].FollowCount = poUser.FollowerCount
		(*dest)[i].FollowerCount = poUser.FollowerCount
		_, boool := followsMap[poUser.ID]
		(*dest)[i].IsFollow = boool
	}
	return nil
}

// GetUserBO 		获取单个用户BO实例
// src				用户PO
// dest				用户BO
func GetUserBO(src *po.User, dest *bo.User) error {
	//先处理isFollow
	//todo暂改
	if middleware.ThreadLocal.Get() == nil { //判断是否登录
		(*dest).IsFollow = false //未登录直接false
	} else { //登录再查询判断
		userId := middleware.ThreadLocal.Get().(map[string]string)[config.Config.ThreadLocal.Keys.UserId]
		uid, err := strconv.Atoi(userId)
		if err != nil {
			return err
		}
		//todo单元测试暂改
		//uid := 1 //单元测试用
		var poFollow = po.Follow{
			FollowId:   (*src).ID,
			FollowerId: uid,
		}
		poFollows, err := daoimpl.NewRelationDaoInstance().QueryByCondition(&poFollow)
		if err != nil {
			return err
		}
		if len(*poFollows) == 0 {
			(*dest).IsFollow = false
		} else {
			(*dest).IsFollow = true
		}
	}
	//再处理其他简单的处理
	dest.ID = src.ID
	dest.Name = src.Name
	dest.FollowCount = src.FollowCount
	dest.FollowerCount = src.FollowerCount
	return nil
}

// GetFeedBOS 	将FeedPo集合转化为FeedBo集合
// src			FeedPO集合
// dest			FeedBO集合
func GetFeedBOS(src *[]po.Feed, dest *[]bo.Feed) {
	if *dest == nil || cap(*dest) < len(*src) {
		temp := make([]bo.Feed, len(*src))
		dest = &temp
	}
	for index, feed := range *src {
		(*dest)[index] = bo.Feed{VideoId: feed.VideoId, CreateTime: feed.CreateTime}
	}
}
