# go-rabbitmq - Wrapper for rabbitmq

## Connect to rabbitmq

```go
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
```

## Create a Queue

```go
    q, err := rbt.QueueDeclare("hello", false, false, false, false, nil)
    utils.FailOnError(err, "Failed to declare queue")
```

## Publish a Message

```go
    err := rbt.Publish("exchange","subject", []byte("Hello World!"))
    utils.FailOnError(err, "Failed to publish message")
```

More examples can be found at [examples](examples/).
## Versioning and license

We use SemVer for versioning. You can see the versions available by checking the tags on this repository.

For more details about our license model, please take a look at the LICENSE file

---

2022, thiagozs
