package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/thiagozs/go-rabbitmq/pkg/utils"
	"github.com/thiagozs/go-rabbitmq/rmq"
)

func main() {

	opts := []rmq.Options{
		rmq.Url("amqp://guest:guest@localhost:5672"),
		rmq.Durable(false),
		rmq.AutoAck(true),
		rmq.Exclusive(false),
		rmq.NoLocal(false),
		rmq.NoWait(false),
		rmq.Args(nil),
		rmq.Name("hello"),
	}

	rbt, err := rmq.NewService(opts...)
	utils.FailOnError(err, "Failed to connect to RabbitMQ")
	defer rbt.Close()

	queue, err := rbt.QueueDeclareWithOpts()
	utils.FailOnError(err, "Failed to declare a queue")

	msg, err := rbt.ConsumeWithOpts(queue)
	utils.FailOnError(err, "Failed to register a consumer")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

	for {
		select {
		case d := <-msg:
			if d.Body == nil {
				<-time.After(time.Second * 1)
				log.Printf("Wait incomming message: %s", d.Body)
				continue
			}
			log.Printf("Received a message: %s", d.Body)
		case <-quit:
			log.Printf("Shutting down")
			return

		case <-rbt.ConnectionDown():
			log.Printf("Connection down")
			queue, err = rbt.QueueDeclareWithOpts()
			utils.FailOnError(err, "Failed to declare a queue")

			msg, err = rbt.ConsumeWithOpts(queue)
			utils.FailOnError(err, "Failed to register a consumer")
		}
	}
}
