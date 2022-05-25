// Package rabbitutil
// @Author shaofan
// @Date 2022/5/15
package rabbitutil

import (
	"douyin/config"
	"github.com/streadway/amqp"
	"log"
)

var (
	conn    *amqp.Connection
	channel *amqp.Channel
)

// Init rabbitmq初始化
func Init() {
	// 建立连接
	var err error
	conn, err = amqp.Dial(config.Config.Rabbit.Url)
	if err != nil {
		log.Fatalln(err)
	}
	channel, err = conn.Channel()
	if err != nil {
		log.Fatalln(err)
	}
	err = channel.Confirm(false)
	if err != nil {
		log.Fatalln(err)
	}
	// 开始监听消费
	initConsumer()
}

// 声明交换机
func initExchange(exchange string) error {
	err := channel.ExchangeDeclare(exchange, "direct", true, false, false, false, nil)
	if err != nil {
		return err
	}
	return nil
}

// 声明队列
func initQueue(queue string) error {
	_, err := channel.QueueDeclare(queue, true, false, false, false, nil)
	if err != nil {
		return err
	}
	return nil
}

// 声明绑定
func initBinding(exchange, queue, key string) error {
	err := channel.QueueBind(queue, key, exchange, false, nil)
	if err != nil {
		return err
	}
	return nil
}

// 生产者发消息之前初始化队列和绑定
func producerInit(exchange, queue, key string) error {
	// 声明交换机
	err := initExchange(config.Config.Rabbit.Exchange.ServiceExchange)
	if err != nil {
		return err
	}
	// 声明队列
	err = initQueue(config.Config.Rabbit.Queue.FeedVideo)
	if err != nil {
		return err
	}
	// 声明绑定
	err = initBinding(
		exchange,
		queue,
		key,
	)
	if err != nil {
		return err
	}
	return nil
}
