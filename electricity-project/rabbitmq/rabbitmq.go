package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

const MQ_URL = "amqp://test:test@106.14.5.37:8088/test"

type RabbitMQ struct {
	conn *amqp.Connection
	channel *amqp.Channel
	// 队列名称
	QueueName string
	// 交换机
	Exchange string
	// key
	Key string
	// 链接信息
	Mqurl string
}

func NewRabbitMQ(queueName, exchange, key string) *RabbitMQ {
	rabbitmq := &RabbitMQ{QueueName:queueName, Exchange:exchange, Key:key, Mqurl:MQ_URL}
	var err error

	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnError(err, "创建链接错误")
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnError(err, "获取channel失败")

	return rabbitmq
}

// 断开channel和connection
// **给结构体定义函数
func (r *RabbitMQ) Destroy() {
	r.channel.Close()
	r.conn.Close()
}

// 错误处理
func (r *RabbitMQ) failOnError(err error, message string) {
	if err != nil {
		log.Fatalf("%s:%s", message, err)
		panic(fmt.Sprintf("%s:%s", message, err))
	}
}

// 创建简单模式下rabbitmq
func NewRabbitMQSimple(queueName string) *RabbitMQ {
	return NewRabbitMQ(queueName, "", "")
}

func (r *RabbitMQ) PublishSimple(message string) {
	// 1.申请队列，如果队列不存在会自动创建，如果存在测跳过创建
	_, err := r.channel.QueueDeclare(
		r.QueueName, false, false, false, false, nil)

	if err != nil {
		fmt.Println(err)
	}

	// 2.发送消息到队列中
	r.channel.Publish(
		r.Exchange, r.Key, false, false,
		amqp.Publishing{ContentType:"text/plain", Body:[]byte(message)})
}
