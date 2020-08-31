package main

import (
	"fmt"
	"go-practice/electricity-project/rabbitmq"
)

func main()  {
	rabbitmq := rabbitmq.NewRabbitMQSimple("test")
	fmt.Sprintf(rabbitmq.Mqurl)
}
