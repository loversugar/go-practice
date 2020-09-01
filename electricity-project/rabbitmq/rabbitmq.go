package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

const MQ_URL = "amqp://test1:test1@xx.xx.xx.xx:8088/test1"

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

//简单模式下生产代码
func (r *RabbitMQ) PublishSimple(message string) {
	//1.申请队列,如果队列不存在会自动创建,如果存在则跳过创建
	//保证队列存在,消息队列能发送到队列中
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		//是否持久化
		false,
		//是否为自动删除
		false,
		//是否具有排他性
		false,
		//是否阻塞
		false,
		//额外属性
		nil,
	)
	if err != nil {
		fmt.Println("QueueDeclare:", err)
	}

	//2.发送消息到队列中
	err = r.channel.Publish(
		r.Exchange,
		r.QueueName,
		//如果为true,根据exchange类型和routekey规则,如果无法找到符合条件的队列那么会把发送的消息返回给发送者
		false,
		//如果为true,当exchange发送消息队列到队列后发现队列上没有绑定消费者,则会把消息发还给发送者
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	if err != nil {
		fmt.Println("Publish:", err)
	}
}

func (r *RabbitMQ) ConsumeSimple()  {
	// 1.申请队列，如果队列不存在会自动创建，如果存在测跳过创建
	_, err := r.channel.QueueDeclare(
		r.QueueName, false, false, false, false, nil)

	if err != nil {
		fmt.Println(err)
	}

	msgs, err := r.channel.Consume(
		r.QueueName, "", true, false, false, false, nil)

	if err != nil {
		fmt.Println(err)
	}

	forever := make(chan bool)
	// 启用协程处理消息
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Println("waiting for messages, to Exit press CTRL+C")
	<- forever
}
