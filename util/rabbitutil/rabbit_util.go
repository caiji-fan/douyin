// Package rabbitutil
// @Author shaofan
// @Date 2022/5/13
// @DESC rabbitmq连接初始化与工具
package rabbitutil

import (
	"douyin/config"
	"douyin/entity/bo"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
)

const MANDATORY = true

// ChangeFollowNum 		修改用户粉丝数和关注数
// userId 				发起关注或取关的用户id
// toUserId 			收到关注或取关的用户id
// isFollow 			是否是关注请求
func ChangeFollowNum(userId int, toUserId int, isFollow bool) error {
	// 服务端声明
	err := producerInit(
		config.Config.Rabbit.Exchange.ServiceExchange,
		config.Config.Rabbit.Queue.UploadVideo,
		config.Config.Rabbit.Key.UploadVideo,
	)
	if err != nil {
		return err
	}
	var body = struct {
		UserId   int  `json:"user_id"`
		ToUserId int  `json:"to_user_id"`
		IsFollow bool `json:"is_follow"`
	}{UserId: userId, ToUserId: toUserId, IsFollow: isFollow}
	fmt.Println(body)
	// 创建消息与管道
	rabbitMSG := bo.RabbitMSG{Data: body, ResendCount: 0}

	data, err := json.Marshal(rabbitMSG)
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	fmt.Println(rabbitMSG)
	channel, err := conn.Channel()
	if err != nil {
		return err
	}
	// 发送消息
	err = channel.Publish(
		config.Config.Rabbit.Exchange.ServiceExchange,
		config.Config.Rabbit.Key.ChangeFollowNum,
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

// UploadVideo 			上传视频文件
// filePath 			视频文件路径
func UploadVideo(filePath string) error {
	// 服务端声明
	err := producerInit(
		config.Config.Rabbit.Exchange.ServiceExchange,
		config.Config.Rabbit.Queue.UploadVideo,
		config.Config.Rabbit.Key.UploadVideo,
	)
	if err != nil {
		return err
	}
	// 创建消息与管道
	rabbitMSG := bo.RabbitMSG{Data: filePath, ResendCount: 0}
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
		config.Config.Rabbit.Exchange.ServiceExchange,
		config.Config.Rabbit.Key.FeedVideo,
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

// FeedVideo 			投放视频到用户feed流
// videoId 				视频id
func FeedVideo(videoId int) error {
	// 声明服务端
	err := producerInit(
		config.Config.Rabbit.Exchange.ServiceExchange,
		config.Config.Rabbit.Queue.FeedVideo,
		config.Config.Rabbit.Key.FeedVideo,
	)
	if err != nil {
		return err
	}
	// 创建消息与管道
	rabbitMSG := bo.RabbitMSG{Data: videoId, ResendCount: 0}
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
		config.Config.Rabbit.Exchange.ServiceExchange,
		config.Config.Rabbit.Key.FeedVideo,
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
