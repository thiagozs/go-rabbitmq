package main

import (
	"fmt"
	"log"

	"github.com/thiagozs/go-rabbitmq/pkg/utils"
	"github.com/thiagozs/go-rabbitmq/rmq"
)

func main() {

	opts := []rmq.Options{
		rmq.Url("amqp://root:secret@localhost:5672/dev"),
		rmq.Durable(false),
		rmq.AutoAck(true),
		rmq.Exclusive(false),
		rmq.NoLocal(false),
		rmq.NoWait(false),
		rmq.Mandatory(false),
		rmq.Immediate(false),
		rmq.Args(nil),
		rmq.Name("hello"),
		rmq.Exchange("teste"),
		rmq.Kind("fanout"),
	}

	rbt, err := rmq.NewService(opts...)
	utils.FailOnError(err, "Failed to connect to RabbitMQ")
	defer rbt.Close()

	err = rbt.ExchangeDeclareWithOpts()
	utils.FailOnError(err, "Failed to declare a queue")

	for i := 0; i < 10; i++ {
		message := fmt.Sprintf("Hello World %d", i)
		log.Printf("Send message: %s", message)
		err = rbt.PublishWithOpts([]byte(message))
		utils.FailOnError(err, "Failed to publish a message")
	}

}
