// Package job
// @Author shaofan
// @Date 2022/5/14
package job

import (
	"douyin/config"
	"douyin/entity/bo"
	"douyin/util/rabbitutil"
	"douyin/util/redisutil"
	"github.com/go-redis/redis"
	"github.com/robfig/cron/v3"
	"log"
	"time"
)

// StartJob 开启任务调度
func StartJob() {
	clearOutBoxJob()
	clearLocalVideo()
	handleErrorMSGJob()
}

// 清理用户发件箱
func clearOutBoxJob() {
	c := cron.New()
	_, err := c.AddFunc("@every 24h", clearOutBox)
	if err != nil {
		log.Fatalln(err)
	}
	c.Start()
}

// 处理错误消息
func handleErrorMSGJob() {
	c := cron.New()
	_, err := c.AddFunc("@every 10s", handleErrorMSG)
	if err != nil {
		log.Fatalln(err)
		return
	}
	c.Start()
}

// todo wangyingsong 清理本地临时存储的视频
func clearLocalVideo() {

}

// 清理用户发件箱
func clearOutBox() {
	var outBoxes = make([]string, 0)
	err := redisutil.Keys(config.Config.Redis.Key.Outbox, &outBoxes)
	if err != nil {
		log.Println(err)
	}
	for _, key := range outBoxes {
		var feeds = make([]bo.Feed, 0)
		expireTime, err := redisutil.TTL(key)
		if err != nil {
			log.Println(err)
		}
		err = redisutil.ZRevRange[bo.Feed](key, &feeds)
		if err != nil {
			log.Println(err)
		}
		var index int
		var feed bo.Feed
		for index, feed = range feeds {
			// 如果发布时间增加七天大于当前，则退出循环
			if feed.CreateTime.AddDate(0, 0, 7).After(time.Now()) {
				break
			}
		}
		// 根据index来获取后面在七天内的数据
		feeds = feeds[index:]
		var value = make([]redis.Z, len(feeds))
		for i, v := range feeds {
			value[i] = redis.Z{
				Score:  float64(v.CreateTime.Unix()),
				Member: v,
			}
		}
		err = redisutil.ZAddWithExpireTime(key, value, expireTime, false, nil)
		if err != nil {
			log.Println(err)
		}
	}
}

// 处理错误消息
func handleErrorMSG() {
	var msgS = make([]bo.RabbitMSG[interface{}], 0)
	err := redisutil.GetAndDelete[[]bo.RabbitMSG[interface{}]](config.Config.Redis.Key.ErrorMessage, &msgS)
	if err != nil {
		log.Println(err)
	}
	for _, msg := range msgS {
		switch msg.Type {
		case bo.FEED_VIDEO:
			err := rabbitutil.Publish[int](&bo.RabbitMSG[int]{Type: msg.Type, Data: msg.Data.(int), ResendCount: msg.ResendCount},
				config.Config.Rabbit.Exchange.ServiceExchange,
				config.Config.Rabbit.Key.FeedVideo)
			if err != nil {
				log.Println(err)
			}
		case bo.UPLOAD_VIDEO:
			err := rabbitutil.Publish[int](&bo.RabbitMSG[int]{Type: msg.Type, Data: msg.Data.(int), ResendCount: msg.ResendCount},
				config.Config.Rabbit.Exchange.ServiceExchange,
				config.Config.Rabbit.Key.UploadVideo)
			if err != nil {
				log.Println(err)
			}
		case bo.CHANGE_FOLLOW_NUM:
			err := rabbitutil.Publish[rabbitutil.ChangeFollowNumBody](
				&bo.RabbitMSG[rabbitutil.ChangeFollowNumBody]{Type: msg.Type, Data: msg.Data.(rabbitutil.ChangeFollowNumBody), ResendCount: msg.ResendCount},
				config.Config.Rabbit.Exchange.ServiceExchange,
				config.Config.Rabbit.Key.ChangeFollowNum)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
