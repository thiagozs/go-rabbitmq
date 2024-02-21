//go:generate ifacemaker -f $GOFILE -s RabbitMQ -i RabbitMQRepo -p rmq -o rmq_repo.go
//go:generate mockgen -source=rmq_repo.go -package=${GOPACKAGE} -destination=${GOPACKAGE}_mock.go
//go:generate mockgen -source=chan_repo.go -package=${GOPACKAGE} -destination=chan_mock.go
//go:generate mockgen -source=conn_repo.go -package=${GOPACKAGE} -destination=conn_mock.go

package rmq

import (
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type GetConnection func() (ConnectionRepo, ChannelRepo, error)

type RabbitMQ struct {
	conn      ConnectionRepo
	ch        ChannelRepo
	params    *RmqParams
	quitM     chan bool
	connCheck chan bool
}

func NewService(opts ...Options) (*RabbitMQ, error) {
	rmq := &RabbitMQ{
		quitM:     make(chan bool),
		connCheck: make(chan bool),
	}

	params, err := newRmqParams(opts...)
	if err != nil {
		return nil, err
	}

	rmq.params = params

	if rmq.params.GetMock() != nil {
		rmq.conn = rmq.params.GetMock().GetConn()
		rmq.ch = rmq.params.GetMock().GetCh()
		return rmq, nil
	}

	conn, ch, err := rmq.getConnection()
	if err != nil {
		return nil, err
	}

	rmq.conn = conn
	rmq.ch = ch

	go rmq.monitoringConn(rmq.getConnection)

	return rmq, nil

}

func (r *RabbitMQ) getConnection() (ConnectionRepo, ChannelRepo, error) {

	if r.params.GetMock() != nil {
		return r.params.GetMock().GetConn(),
			r.params.GetMock().GetCh(), nil
	}

	conn, err := amqp.Dial(r.params.GetUrl())
	if err != nil {
		return nil, nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, nil, err
	}

	return conn, ch, nil
}

func (r *RabbitMQ) monitoringConn(cb GetConnection) {
	go func(r *RabbitMQ) {
		var heartBeat time.Duration

		if r.params.GetHeartbeat() == "" {
			heartBeat, _ = time.ParseDuration("200ms")
		}

		for {
			select {
			case <-time.After(heartBeat):
				if !r.conn.IsClosed() {
					continue
				}

				conn, ch, err := cb()
				if err != nil {
					continue
				}

				r.conn = conn
				r.ch = ch

				r.connCheck <- true

			case <-r.quitM:
				return
			}
		}
	}(r)
}

func (r *RabbitMQ) Close() error {
	if err := r.ch.Close(); err != nil {
		return err
	}

	if err := r.conn.Close(); err != nil {
		return err
	}

	select {
	case r.quitM <- true:
		// Sent successfully
		fmt.Println("Close monitoring connection, receiver ready")
	default:
		// No receiver ready, proceed without blocking
		fmt.Println("Close monitoring connection, no receiver")
	}

	return nil
}

func (r *RabbitMQ) Publish(exchange, routingKey string, body []byte) error {
	return r.GetCh().Publish(exchange, routingKey, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        body,
	})
}

func (r *RabbitMQ) Consume(queue, consumer string, autoAck,
	exclusive, noLocal, noWait bool, args amqp.Table) (<-chan amqp.Delivery, error) {
	return r.GetCh().Consume(queue, consumer, autoAck, exclusive, noLocal, noWait, args)
}

func (r *RabbitMQ) QueueDeclare(name string, durable, autoDelete,
	exclusive, noWait bool, args amqp.Table) (amqp.Queue, error) {
	return r.GetCh().QueueDeclare(name, durable, autoDelete, exclusive, noWait, args)
}

func (r *RabbitMQ) ExchangeDeclare(name, kind string, durable,
	autoDelete, internal, noWait bool, args amqp.Table) error {
	return r.GetCh().ExchangeDeclare(name, kind, durable, autoDelete, internal, noWait, args)
}

func (r *RabbitMQ) QueueBind(name, key, exchange string,
	noWait bool, args amqp.Table) error {
	return r.GetCh().QueueBind(name, key, exchange, noWait, args)
}

func (r *RabbitMQ) QueueUnbind(name, key, exchange string, args amqp.Table) error {
	return r.GetCh().QueueUnbind(name, key, exchange, args)
}

func (r *RabbitMQ) QueueDelete(name string, ifUnused, ifEmpty, noWait bool) (int, error) {
	return r.GetCh().QueueDelete(name, ifUnused, ifEmpty, noWait)
}

func (r *RabbitMQ) ExchangeDelete(name string, ifUnused, noWait bool) error {
	return r.GetCh().ExchangeDelete(name, ifUnused, noWait)
}

func (r *RabbitMQ) PublishWithOpts(body []byte) error {
	return r.GetCh().Publish(r.params.GetExchange(),
		r.params.GetKey(), r.params.GetMandatory(),
		false, amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
}

func (r *RabbitMQ) ConsumeWithOpts() (<-chan amqp.Delivery, error) {
	return r.GetCh().Consume(r.params.GetName(), r.params.GetConsumer(),
		r.params.GetAutoAck(), r.params.GetExclusive(),
		r.params.GetNoLocal(), r.params.GetNoWait(),
		r.params.GetArgs())
}

func (r *RabbitMQ) QueueDeclareWithOpts() (amqp.Queue, error) {
	q, err := r.ch.QueueDeclare(r.params.GetName(), r.params.GetDurable(),
		r.params.GetAutoDelete(), r.params.GetExclusive(),
		r.params.GetNoWait(), r.params.GetArgs())
	if err != nil {
		return q, err
	}

	if err := r.QueueBind(q.Name, r.params.GetKey(), r.params.GetExchange(),
		r.params.GetNoWait(), r.params.GetArgs()); err != nil {
		return q, err
	}

	return q, nil
}

func (r *RabbitMQ) ExchangeDeclareWithOpts() error {
	return r.GetCh().ExchangeDeclare(r.params.GetExchange(),
		r.params.GetKind(), r.params.GetDurable(),
		r.params.GetAutoDelete(), r.params.GetInternal(),
		r.params.GetNoWait(), r.params.GetArgs())
}

func (r *RabbitMQ) IsDisconected() bool {
	return r.conn.IsClosed()
}

func (r *RabbitMQ) MonitorConn() <-chan bool {
	return r.connCheck
}

func (r *RabbitMQ) GetConn() ConnectionRepo {
	return r.conn
}

func (r *RabbitMQ) GetCh() ChannelRepo {
	return r.ch
}
