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
	"os"
	"time"
)

// StartJob 开启任务调度
func StartJob() {
	clearOutBoxJob()
	clearLocalVideoJob()
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

// 清理本地临时存储的视频
func clearLocalVideoJob() {
	c := cron.New()
	_, err := c.AddFunc("@every 168h", clearLocalVideo)
	if err != nil {
		log.Fatalln(err)
		return
	}
	c.Start()
}

// 清理用户发件箱
func clearOutBox() {
	var outBoxes = make([]string, 0)
	err := redisutil.Keys(config.Config.Redis.Key.Outbox, &outBoxes)
	if err != nil {
		log.Println(err)
		return
	}
	for _, key := range outBoxes {
		var feeds = make([]bo.Feed, 0)
		if err != nil {
			log.Println(err)
			return
		}
		// 获取发件箱中的数据
		err = redisutil.ZRevRange[bo.Feed](key, &feeds)
		if err != nil {
			log.Println(err)
			return
		}
		var index int
		var feed bo.Feed
		// 对发件箱中的数据进行排查
		for index, feed = range feeds {
			// 如果发布时间增加七天大于当前，则退出循环
			if feed.CreateTime.AddDate(0, 0, 7).After(time.Now()) {
				break
			}
		}
		// 根据index来获取需要清理的数据
		feeds = feeds[:index]
		err = redisutil.ZRem[bo.Feed](key, &feeds, false, nil)
		if err != nil {
			log.Println(err)
			return
		}
	}
}

// 处理错误消息
func handleErrorMSG() {
	var msgS = make([]bo.RabbitMSG[interface{}], 0)
	err := redisutil.GetAndDelete[[]bo.RabbitMSG[interface{}]](config.Config.Redis.Key.ErrorMessage, &msgS)
	if err != nil {
		log.Println(err)
		return
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

// 清理本地临时存储的视频
func clearLocalVideo() {
	// 获得两周前的日期
	now := time.Now().AddDate(0, 0, -14).Format(config.Config.StandardDate)
	// 拼接两周前的视频和文件路径
	videoPath := config.Config.Service.VideoTempDir + now
	coverPath := config.Config.Service.CoverTempDir + now
	// 判断视频路径是否存在并删除
	_, err := os.Stat(videoPath)
	if err != nil {
		log.Println(err)
		return
	}
	err = os.RemoveAll(videoPath)
	if err != nil {
		log.Println(err)
		return
	}
	// 判断封面路径是否存在并删除
	_, err = os.Stat(coverPath)
	if err != nil {
		log.Println(err)
		return
	}
	err = os.RemoveAll(coverPath)
	if err != nil {
		log.Println(err)
		return
	}
}
