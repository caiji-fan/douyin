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
)

// StartJob 开启任务调度
func StartJob() {
	clearOutBox()
	clearInBox()
	clearLocalVideo()
	handleErrorMSG()
}

// 清理用户收件箱
func clearInBox() {

}

// 清理用户发件箱
func clearOutBox() {

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
	}
	c.Start()
}

// 清理本地临时存储的视频
func clearLocalVideo() {

}
