package main

import (
	"fmt"
	"log"

	"github.com/thiagozs/go-rabbitmq/pkg/utils"
	"github.com/thiagozs/go-rabbitmq/rmq"
)

func main() {

	opts := []rmq.Options{
		rmq.WithUrl("amqp://guest:guest@localhost:5672"),
		rmq.WithDurable(true),
		rmq.WithAutoAck(true),
		rmq.WithExclusive(false),
		rmq.WithNoLocal(false),
		rmq.WithNoWait(false),
		rmq.WithMandatory(false),
		rmq.WithImmediate(false),
		rmq.WithArgs(nil),
		rmq.WithName("RoutingProcessed"),
		rmq.WithExchange("ASAPExchange.Authorizer"),
		rmq.WithKind("direct"),
	}

	rbt, err := rmq.NewService(opts...)
	utils.FailOnError(err, "Failed to connect to RabbitMQ")
	defer rbt.Close()

	err = rbt.ExchangeDeclareWithOpts()
	utils.FailOnError(err, "Failed to declare a queue")

	for i := 0; i < 1000; i++ {
		message := fmt.Sprintf("Hello World %d", i)
		log.Printf("Send message: %s", message)
		err = rbt.PublishWithOpts([]byte(message))
		utils.FailOnError(err, "Failed to publish a message")
	}

}
