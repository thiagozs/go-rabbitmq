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
		rmq.WithUrl("amqp://guest:guest@localhost:5672"),
		rmq.WithDurable(true),
		rmq.WithAutoAck(true),
		rmq.WithExclusive(false),
		rmq.WithNoLocal(false),
		rmq.WithNoWait(false),
		rmq.WithMandatory(false),
		rmq.WithImmediate(false),
		rmq.WithArgs(nil),
		rmq.WithName("collector_transaction_queue"),
		rmq.WithExchange("ASAPExchange.Authorizer"),
		rmq.WithKind("direct"),
		rmq.WithKey("RoutingProcessed"),
	}

	rbt, err := rmq.NewService(opts...)
	utils.FailOnError(err, "Failed to connect to RabbitMQ")
	defer rbt.Close()

	_, err = rbt.QueueDeclareWithOpts()
	utils.FailOnError(err, "Failed to declare a queue")

	msg, err := rbt.ConsumeWithOpts()
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
			if err := rbt.Close(); err != nil {
				log.Printf("Failed to close connection")
			}
			return

		case <-rbt.MonitorConn():
			log.Printf("Connection down")
			_, err = rbt.QueueDeclareWithOpts()
			utils.FailOnError(err, "Failed to declare a queue")

			msg, err = rbt.ConsumeWithOpts()
			utils.FailOnError(err, "Failed to register a consumer")
			log.Printf("Connection recovered")
		}
	}
}
