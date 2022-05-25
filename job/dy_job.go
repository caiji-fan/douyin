// Package job
// @Author shaofan
// @Date 2022/5/14
package job

import (
	"douyin/config"
	"douyin/entity/bo"
	"douyin/util/rabbitutil"
	"douyin/util/redisutil"
	"github.com/robfig/cron/v3"
	"log"
	"time"
)

// StartJob 开启任务调度
func StartJob() {
	clearOutBox()
	clearLocalVideo()
	handleErrorMSG()
}

// 清理用户发件箱
func clearOutBox() {
	c := cron.New()
	_, err := c.AddFunc("@every 24h", func() {
		var outBoxes = make([]string, 0)
		err := redisutil.Keys(config.Config.Redis.Key.Outbox, outBoxes)
		if err != nil {
			log.Println(err)
		}
		for _, key := range outBoxes {
			var feeds = make([]bo.Feed, 0)
			expireTime, err := redisutil.GetExpireTime(key)
			if err != nil {
				log.Println(err)
			}
			err = redisutil.ZGet(key, feeds)
			if err != nil {
				log.Println(err)
			}
			var index int
			var feed bo.Feed
			for index, feed = range feeds {
				createTime, err := time.Parse(config.Config.StandardTime, feed.CreateTime)
				if err != nil {
					log.Println(err)
				}
				// 如果发布时间增加七天大于当前，则退出循环
				if createTime.AddDate(0, 0, 7).After(time.Now()) {
					break
				}
			}
			// 根据index来获取后面在七天内的数据
			feeds = feeds[index:]
			err = redisutil.ZSetWithExpireTime(key, feeds, "CreateTime", expireTime)
			if err != nil {
				log.Println(err)
			}
		}
	})
	if err != nil {
		log.Fatalln(err)
	}
	c.Start()
}

// 处理错误消息
func handleErrorMSG() {
	c := cron.New()
	_, err := c.AddFunc("@every 10s", func() {
		var msgS = make([]bo.RabbitMSG, 0)
		err := redisutil.GetAndDelete(config.Config.Redis.Key.ErrorMessage, &msgS)
		if err != nil {
			log.Println(err)
		}
		for _, msg := range msgS {
			switch msg.Type {
			case bo.FEED_VIDEO:
				err := rabbitutil.Publish(&msg, config.Config.Rabbit.Exchange.ServiceExchange, config.Config.Rabbit.Key.FeedVideo)
				if err != nil {
					log.Println(err)
				}
			case bo.UPLOAD_VIDEO:
				err := rabbitutil.Publish(&msg, config.Config.Rabbit.Exchange.ServiceExchange, config.Config.Rabbit.Key.UploadVideo)
				if err != nil {
					log.Println(err)
				}
			case bo.CHANGE_FOLLOW_NUM:
				err := rabbitutil.Publish(&msg, config.Config.Rabbit.Exchange.ServiceExchange, config.Config.Rabbit.Key.ChangeFollowNum)
				if err != nil {
					log.Println(err)
				}
			}
		}
	})
	if err != nil {
		log.Fatalln(err)
		return
	}
	c.Start()
}

// 清理本地临时存储的视频
func clearLocalVideo() {

}
