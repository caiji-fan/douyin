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
	"github.com/go-redis/redis"
	"log"
	"strconv"
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
			var rabbitMSG bo.RabbitMSG[ChangeFollowNumBody]
			//反序列化
			err := json.Unmarshal(msg.Body, &rabbitMSG)
			failOnError[ChangeFollowNumBody](err, &rabbitMSG)
			changeFollowNumBody := rabbitMSG.Data
			err = doChangeFollowNum(&changeFollowNumBody)
			failOnError[ChangeFollowNumBody](err, &rabbitMSG)
			err = msg.Ack(true)
			failOnError[ChangeFollowNumBody](err, &rabbitMSG)
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
			var rabbitMSG bo.RabbitMSG[int]
			//反序列化
			err := json.Unmarshal(msg.Body, &rabbitMSG)
			failOnError[int](err, &rabbitMSG)
			//查询video数据
			videoId := rabbitMSG.Data
			err = doUploadVideo(videoId)
			failOnError[int](err, &rabbitMSG)
			// 确认收到消息
			err = msg.Ack(true)
			failOnError[int](err, &rabbitMSG)
		}
	}()
}

//错误处理
func failOnError[T any](err error, rabbitMSG *bo.RabbitMSG[T]) {
	if err != nil {
		if int(rabbitMSG.ResendCount) > config.Config.Rabbit.ResendMax {
			// todo 报警
		}
		handleError[T](rabbitMSG)
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
			var rabbitMSG bo.RabbitMSG[int]
			//反序列化
			err := json.Unmarshal(msg.Body, &rabbitMSG)
			failOnError[int](err, &rabbitMSG)
			//查询video数据
			videoId := rabbitMSG.Data
			err = doFeedVideo(videoId)
			failOnError[int](err, &rabbitMSG)
			err = msg.Ack(true)
			failOnError[int](err, &rabbitMSG)
		}
	}()
}

// 消息补偿机制
func handleError[T any](msg *bo.RabbitMSG[T]) {
	var rabbitMSGS = make([]bo.RabbitMSG[T], 0)
	err := redisutil.Get[[]bo.RabbitMSG[T]](config.Config.Redis.Key.ErrorMessage, &rabbitMSGS)
	failOnError[T](err, msg)
	rabbitMSGS = append(rabbitMSGS, *msg)
	err = redisutil.Set(config.Config.Redis.Key.ErrorMessage, &rabbitMSGS)
	failOnError[T](err, msg)
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
		difference = 1
	} else {
		difference = -1
	}
	// 增减发起请求的一方
	err = daoimpl.NewUserDaoInstance().ChangeFollowCount(body.UserId, difference, tx, true)
	if err != nil {
		tx.Rollback()
		return err
	}
	// 增减接收请求的一方
	err = daoimpl.NewUserDaoInstance().ChangeFansCount(body.ToUserId, difference, tx, true)
	if err != nil {
		tx.Rollback()
		return err
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
	// 如果是视频不存在，结束处理
	if video.ID == 0 {
		return nil
	}
	//上传视频
	videoUrl, err := obsutil.Upload(video.PlayUrl, config.Config.Obs.Buckets.Video)
	video.PlayUrl = videoUrl
	if err != nil {
		return err
	}
	// 上传封面
	coverUrl, err := obsutil.Upload(video.CoverUrl, config.Config.Obs.Buckets.Cover)
	if err != nil {
		return err
	}
	video.CoverUrl = coverUrl

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
	// 如果是视频不存在，结束处理
	if video.ID == 0 {
		return nil
	}
	sender, err := daoimpl.NewUserDaoInstance().QueryById(video.AuthorId)
	if err != nil {
		return err
	}
	var videos []bo.Feed
	//大v用户
	if sender.FollowerCount >= config.Config.Service.BigVNum {
		err = redisutil.ZRevRange[bo.Feed](config.Config.Redis.Key.Outbox+strconv.Itoa(sender.ID), &videos)
		if err != nil {
			return err
		}
		videos = append(videos, bo.Feed{VideoId: videoId, CreateTime: video.CreateTime})
		var value = make([]redis.Z, len(videos))
		for i, v := range videos {
			value[i] = redis.Z{
				Score:  float64(v.CreateTime.UnixMilli()),
				Member: v,
			}
		}
		err = redisutil.ZAddWithExpireTime(config.Config.Redis.Key.Outbox+strconv.Itoa(sender.ID),
			value,
			config.OutboxExpireTime,
			false,
			nil)
		if err != nil {
			return err
		}
	} else if sender.FollowerCount > 0 {
		//普通用户
		userIds, err := daoimpl.NewRelationDaoInstance().QueryFansIdByFollowId(video.AuthorId)
		if err != nil {
			return err
		}
		users, err := daoimpl.NewUserDaoInstance().QueryBatchIds(&userIds)
		if err != nil {
			return err
		}
		// 创建redis事务处理
		var pipeline = redisutil.Begin()
		// feed集合，用户持久化
		var feeds = make([]po.Feed, 0)
		for _, user := range *users {
			err = redisutil.ZRevRange[bo.Feed](config.Config.Redis.Key.Inbox+strconv.Itoa(user.ID), &videos)
			if err != nil {
				return err
			}
			// 不存在收件箱，则入库
			if videos == nil || len(videos) == 0 {
				feeds = append(feeds, po.Feed{UserId: user.ID, VideoId: videoId})
			} else {
				videos = append(videos, bo.Feed{VideoId: videoId, CreateTime: video.CreateTime})
				var value = make([]redis.Z, len(videos))
				for i, v := range videos {
					value[i] = redis.Z{
						Score:  float64(v.CreateTime.UnixMilli()),
						Member: v,
					}
				}
				err = redisutil.ZAddWithExpireTime(config.Config.Redis.Key.Inbox+strconv.Itoa(user.ID),
					value,
					config.InboxExpireTime,
					true,
					pipeline)
				if err != nil {
					err1 := pipeline.Discard()
					if err1 != nil {
						return err1
					}
					return err
				}
			}
		}
		// 如果没有数据入库，则直接执行redis事务后退出
		if len(feeds) == 0 {
			_, err = pipeline.Exec()
			if err != nil {
				return err
			}
			return nil
		}
		// 开始feed持久化
		feedDao := daoimpl.NewFeedDaoInstance()
		tx := feedDao.Begin()
		err = feedDao.InsertBatch(&feeds, tx, true)
		if err != nil {
			tx.Rollback()
			err1 := pipeline.Discard()
			if err1 != nil {
				return err1
			}
			return err
		}
		_, err = pipeline.Exec()
		if err != nil {
			tx.Rollback()
			return err
		}
		tx.Commit()
	}
	return nil
}
