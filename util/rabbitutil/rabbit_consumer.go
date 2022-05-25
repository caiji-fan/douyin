// Package rabbitutil
// @Author shaofan
// @Date 2022/5/15
package rabbitutil

import (
	"douyin/config"
	"douyin/entity/bo"
	"douyin/entity/po"
	"douyin/repositories/daoimpl"
	"douyin/util/obsutil"
	"douyin/util/redisutil"
	"encoding/json"
	"log"
	"strconv"
	"sync"
)

// 初始化consumer
func initConsumer() {
	initRabbitServer()
	changeFollowNumConsumer()
	uploadVideoConsumer()
	feedVideoConsumer()
}

// 修改关注数量消费
func changeFollowNumConsumer() {
	consume, err := channel.Consume(
		config.Config.Rabbit.Queue.ChangeFollowNum,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println(err)
	}
	// 协程处理消费
	go func() {
		for msg := range consume {
			var rabbitMSG bo.RabbitMSG
			//反序列化
			err := json.Unmarshal(msg.Body, &rabbitMSG)
			failOnError(err, &rabbitMSG)
			changeFollowNumBody := rabbitMSG.Data.(ChangeFollowNumBody)
			err = doChangeFollowNum(&changeFollowNumBody)
			failOnError(err, &rabbitMSG)
			err = msg.Ack(true)
			failOnError(err, &rabbitMSG)
		}
	}()
}

// 上传视频消费
func uploadVideoConsumer() {
	consume, err := channel.Consume(
		config.Config.Rabbit.Queue.UploadVideo,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println(err)
	}
	// 协程处理消费
	go func() {
		for msg := range consume {
			var rabbitMSG bo.RabbitMSG
			//反序列化
			err := json.Unmarshal(msg.Body, &rabbitMSG)
			failOnError(err, &rabbitMSG)
			//查询video数据
			videoId, _ := rabbitMSG.Data.(int)
			err = doUploadVideo(videoId)
			failOnError(err, &rabbitMSG)
			// 确认收到消息
			err = msg.Ack(true)
			failOnError(err, &rabbitMSG)
		}
	}()
}

//错误处理
func failOnError(err error, rabbitMSG *bo.RabbitMSG) {
	if err != nil {
		if int(rabbitMSG.ResendCount) > config.Config.Rabbit.ResendMax {
			// todo 报警
		}
		handleError(rabbitMSG)
		log.Println(err)
	}
}

// 投放视频流消费
func feedVideoConsumer() {
	consume, err := channel.Consume(
		config.Config.Rabbit.Queue.FeedVideo,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println(err)
	}
	// 协程处理消费
	go func() {
		for msg := range consume {
			var rabbitMSG bo.RabbitMSG
			//反序列化
			err := json.Unmarshal(msg.Body, &rabbitMSG)
			failOnError(err, &rabbitMSG)
			//查询video数据
			videoId, _ := rabbitMSG.Data.(int)
			err = doFeedVideo(videoId)
			failOnError(err, &rabbitMSG)
			err = msg.Ack(true)
			failOnError(err, &rabbitMSG)
		}
	}()
}

// 消息补偿机制
func handleError(msg *bo.RabbitMSG) {
	var rabbitMSGS = make([]bo.RabbitMSG, 0)
	err := redisutil.Get(config.Config.Redis.Key.ErrorMessage, &rabbitMSGS)
	failOnError(err, msg)
	rabbitMSGS = append(rabbitMSGS, *msg)
	err = redisutil.Set(config.Config.Redis.Key.ErrorMessage, &rabbitMSGS)
	failOnError(err, msg)
}

// rabbit服务器初始化
func initRabbitServer() {
	// 声明交换机
	err := initExchange(config.Config.Rabbit.Exchange.ServiceExchange)
	if err != nil {
		log.Fatalln(err)
	}
	// 声明队列
	err = initQueue(config.Config.Rabbit.Queue.FeedVideo)
	if err != nil {
		log.Fatalln(err)
	}
	// 声明队列
	err = initQueue(config.Config.Rabbit.Queue.UploadVideo)
	if err != nil {
		log.Fatalln(err)
	}
	// 声明队列
	err = initQueue(config.Config.Rabbit.Queue.ChangeFollowNum)
	if err != nil {
		log.Fatalln(err)
	}
	// 声明绑定
	err = initBinding(
		config.Config.Rabbit.Exchange.ServiceExchange,
		config.Config.Rabbit.Queue.FeedVideo,
		config.Config.Rabbit.Key.FeedVideo,
	)
	if err != nil {
		log.Fatalln(err)
	}
	// 声明绑定
	err = initBinding(
		config.Config.Rabbit.Exchange.ServiceExchange,
		config.Config.Rabbit.Queue.UploadVideo,
		config.Config.Rabbit.Key.UploadVideo,
	)
	if err != nil {
		log.Fatalln(err)
	}
	// 声明绑定
	err = initBinding(
		config.Config.Rabbit.Exchange.ServiceExchange,
		config.Config.Rabbit.Queue.ChangeFollowNum,
		config.Config.Rabbit.Key.ChangeFollowNum,
	)
	if err != nil {
		log.Fatalln(err)
	}
}

// 更改关注和粉丝数量
func doChangeFollowNum(body *ChangeFollowNumBody) error {
	var err error
	tx := daoimpl.NewUserDaoInstance().Begin()
	var difference int
	if body.IsFollow {
		difference = -1
	} else {
		difference = 1
	}
	wait := sync.WaitGroup{}
	wait.Add(2)
	//增减发起关注的一方
	go func() {
		defer wait.Done()
		var user *po.User
		user, err = daoimpl.NewUserDaoInstance().QueryForUpdate(body.UserId, tx)
		if err != nil {
			return
		}
		user.FollowCount = user.FollowCount + difference
		err = daoimpl.NewUserDaoInstance().UpdateByCondition(user, tx, true)
	}()
	//增减收到关注的一方
	go func() {
		defer wait.Done()
		var user *po.User
		user, err = daoimpl.NewUserDaoInstance().QueryForUpdate(body.ToUserId, tx)
		if err != nil {
			return
		}
		user.FollowerCount = user.FollowerCount + difference
		err = daoimpl.NewUserDaoInstance().UpdateByCondition(user, tx, true)
	}()
	wait.Wait()
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	return nil
}

// 上传视频
func doUploadVideo(videoId int) error {
	video, err := daoimpl.NewVideoDaoInstance().QueryById(videoId)
	if err != nil {
		return err
	}
	//上传视频
	url, err := obsutil.Upload(video.PlayUrl, config.Config.Obs.Buckets.Video)
	video.PlayUrl = url
	if err != nil {
		return err
	}
	// 更新数据库
	err = daoimpl.NewVideoDaoInstance().UpdateByCondition(video, nil, false)
	if err != nil {
		return err
	}
	return nil
}

// 投放视频流
func doFeedVideo(videoId int) error {
	video, err := daoimpl.NewVideoDaoInstance().QueryById(videoId)
	if err != nil {
		return err
	}
	sender, err := daoimpl.NewUserDaoInstance().QueryById(video.AuthorId)
	if err != nil {
		return err
	}
	// todo redis事务
	var videos []bo.Feed
	//大v用户
	if sender.FollowCount >= config.Config.Service.BigVNum {
		err = redisutil.ZGet(config.Config.Redis.Key.Outbox+strconv.Itoa(sender.ID), &videos)
		if err != nil {
			return err
		}
		videos = append(videos, bo.Feed{VideoId: videoId, CreateTime: video.CreateTime})
		err = redisutil.ZSetWithExpireTime(config.Config.Redis.Key.Outbox+strconv.Itoa(sender.ID),
			&videos,
			"CreateTime",
			config.OutboxExpireTime)
		if err != nil {
			return err
		}
	} else {
		//普通用户
		userIds, err := daoimpl.NewRelationDaoInstance().QueryFansIdByFollowId(video.AuthorId)
		if err != nil {
			return err
		}
		users, err := daoimpl.NewUserDaoInstance().QueryBatchIds(&userIds)
		if err != nil {
			return err
		}
		// feed集合，用户持久化
		// todo 增加redis事务控制
		var feeds = make([]po.Feed, 0)
		for _, user := range *users {
			err = redisutil.ZGet(config.Config.Redis.Key.Inbox+strconv.Itoa(user.ID), &videos)
			if err != nil {
				return err
			}
			// 不存在收件箱，则入库
			if videos == nil {
				feeds = append(feeds, po.Feed{UserId: user.ID, VideoId: videoId})
			} else {
				videos = append(videos, bo.Feed{VideoId: videoId, CreateTime: video.CreateTime})
				err = redisutil.ZSetWithExpireTime(config.Config.Redis.Key.Inbox+strconv.Itoa(user.ID),
					&videos,
					"CreateTime",
					config.InboxExpireTime)
				if err != nil {
					return err
				}
			}
		}
		feedDao := daoimpl.NewFeedDaoInstance()
		err = feedDao.InsertBatch(&feeds)
		if err != nil {
			return err
		}
	}
	return nil
}
