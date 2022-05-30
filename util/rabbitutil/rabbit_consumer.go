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
func initConsumer() error {
	if err := initServer(); err != nil {
		return err
	}
	if err := changeFollowNumConsumer(); err != nil {
		return err
	}
	if err := uploadVideoConsumer(); err != nil {
		return err
	}
	if err := feedVideoConsumer(); err != nil {
		return err
	}
	return nil
}

// 初始化rabbitmq服务器
func initServer() error {
	if err := initFeedVideo(); err != nil {
		return err
	}
	if err := initUploadVideo(); err != nil {
		return err
	}
	if err := initChangeFollowNum(); err != nil {
		return err
	}
	return nil
}

// 修改关注数量消费
func changeFollowNumConsumer() error {
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
		return err
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
	return nil
}

// 上传视频消费
func uploadVideoConsumer() error {
	consume, err := channel.Consume(
		config.Config.Rabbit.Queue.DeadUploadVideo,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
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
	return nil
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
func feedVideoConsumer() error {
	consume, err := channel.Consume(
		config.Config.Rabbit.Queue.DeadFeedVideo,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
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
	return nil
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
	// 获取投稿人的信息
	sender, err := daoimpl.NewUserDaoInstance().QueryById(video.AuthorId)
	if err != nil {
		return err
	}
	// 取得需要存入redis的value
	var value = []redis.Z{{Score: float64(video.CreateTime.UnixMilli()), Member: video}}
	//大v用户
	if sender.FollowerCount >= config.Config.Service.BigVNum {
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
			// 增加到用户的发件箱中,向用户投放不影响用户的收件箱过期时间
			err = redisutil.ZAdd(config.Config.Redis.Key.Inbox+strconv.Itoa(user.ID),
				value,
				true,
				pipeline)
			if err != nil {
				err1 := pipeline.Discard()
				if err1 != nil {
					return err1
				}
				return err
			}
			// 记录需要入库的数据
			feeds = append(feeds, po.Feed{UserId: user.ID, VideoId: videoId})
		}
		// 开始feed持久化，对数据入库
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
