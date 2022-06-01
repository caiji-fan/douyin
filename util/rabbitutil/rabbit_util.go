// Package rabbitutil
// @Author shaofan
// @Date 2022/5/13
// @DESC rabbitmq连接初始化与工具
package rabbitutil

import (
	"douyin/config"
	"douyin/entity/rabbitentity"
	"encoding/json"
	"github.com/streadway/amqp"
)

const MANDATORY = true

// ChangeFollowNum 		修改用户粉丝数和关注数
// userId 				发起关注或取关的用户id
// toUserId 			收到关注或取关的用户id
// isFollow 			是否是关注请求
func ChangeFollowNum(userId int, toUserId int, isFollow bool) error {
	channel, err := conn.Channel()
	if err != nil {
		return err
	}
	// 服务端声明
	if err := initChangeFollowNum(channel); err != nil {
		return err
	}
	var body = rabbitentity.ChangeFollowNumBody{UserId: userId, ToUserId: toUserId, IsFollow: isFollow}
	// 创建消息与管道
	rabbitMSG := rabbitentity.RabbitMSG[rabbitentity.ChangeFollowNumBody]{Data: body, ResendCount: 0, Type: rabbitentity.CHANGE_FOLLOW_NUM}
	return publish[rabbitentity.ChangeFollowNumBody](&rabbitMSG,
		config.Config.Rabbit.Exchange.ServiceExchange,
		config.Config.Rabbit.Key.ChangeFollowNum,
		channel)
}

// UploadVideo 			上传视频文件
// videoId 				视频id
// filePath 			视频文件路径
func UploadVideo(videoId int) error {
	channel, err := conn.Channel()
	if err != nil {
		return err
	}
	if err := initUploadVideo(channel); err != nil {
		return err
	}
	// 创建消息与管道
	rabbitMSG := rabbitentity.RabbitMSG[int]{Data: videoId, ResendCount: 0, Type: rabbitentity.UPLOAD_VIDEO}
	return publish[int](&rabbitMSG,
		config.Config.Rabbit.Exchange.ServiceExchange,
		config.Config.Rabbit.Key.UploadVideo,
		channel)
}

// FeedVideo 			投放视频到用户feed流
// videoId 				视频id
func FeedVideo(videoId int) error {
	channel, err := conn.Channel()
	if err != nil {
		return err
	}
	if err := initFeedVideo(channel); err != nil {
		return err
	}
	// 创建消息与管道
	rabbitMSG := rabbitentity.RabbitMSG[int]{Data: videoId, ResendCount: 0, Type: rabbitentity.FEED_VIDEO}
	return publish[int](&rabbitMSG,
		config.Config.Rabbit.Exchange.ServiceExchange,
		config.Config.Rabbit.Key.FeedVideo,
		channel)
}

// Publish 发布消息，外部使用
func Publish[T rabbitentity.RabbitType](rabbitMSG *rabbitentity.RabbitMSG[T], exchange string, key string) error {
	data, err := json.Marshal(rabbitMSG)
	if err != nil {
		return err
	}
	channel, err := conn.Channel()
	if err != nil {
		return err
	}

	// 发送消息
	err = channel.Publish(
		exchange,
		key,
		MANDATORY,
		false,
		amqp.Publishing{
			ContentType: "text/json",
			Body:        data,
		},
	)
	if err != nil {
		return err
	}
	return channel.Close()
}

//  发布消息内部使用
func publish[T rabbitentity.RabbitType](rabbitMSG *rabbitentity.RabbitMSG[T], exchange string, key string, channel *amqp.Channel) error {
	data, err := json.Marshal(rabbitMSG)
	if err != nil {
		return err
	}

	// 发送消息
	err = channel.Publish(
		exchange,
		key,
		MANDATORY,
		false,
		amqp.Publishing{
			ContentType: "text/json",
			Body:        data,
		},
	)
	if err != nil {
		return err
	}
	err = channel.Close()
	if err != nil {
		return err
	}
	return nil
}
