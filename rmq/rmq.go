package rmq

import (
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type KindGetChannels func(cfg Config) (*amqp.Connection, *amqp.Channel, error)

type RabbitMQ struct {
	Conn     *amqp.Connection
	Ch       *amqp.Channel
	Cfg      Config
	quitM    chan bool
	Opts     []Options
	connDown chan bool
}

func NewService(opts ...Options) (*RabbitMQ, error) {
	cfg := NewOptions(opts...)

	return InitServices(cfg, opts)
}

func InitServices(cfg Config, opts []Options) (*RabbitMQ, error) {

	getChannels := func(cfg Config) (*amqp.Connection, *amqp.Channel, error) {
		conn, err := amqp.Dial(cfg.Url)
		if err != nil {
			return nil, nil, err
		}

		ch, err := conn.Channel()
		if err != nil {
			return nil, nil, err
		}

		return conn, ch, nil
	}

	conn, ch, err := getChannels(cfg)
	if err != nil {
		return nil, err
	}

	rmq := &RabbitMQ{
		Conn:     conn,
		Ch:       ch,
		Cfg:      cfg,
		quitM:    make(chan bool),
		Opts:     opts,
		connDown: make(chan bool),
	}

	go rmq.monitoringConn(getChannels)

	return rmq, nil
}

func (r *RabbitMQ) monitoringConn(fn KindGetChannels) {
	go func(c Config, conn *amqp.Connection) {
		var heartBeat time.Duration

		if c.Heartbeat == "" {
			heartBeat, _ = time.ParseDuration("200ms")
		}

		for {
			select {
			case <-time.After(heartBeat):
				if !conn.IsClosed() {
					continue
				}

				nconn, ch, err := fn(c)
				if err != nil {
					continue
				}

				conn = nconn
				r.Conn = nconn
				r.Ch = ch

				r.connDown <- true

			case <-r.quitM:
				return
			}
		}
	}(r.Cfg, r.Conn)
}

func (r *RabbitMQ) Close() {
	r.Conn.Close()
	r.Ch.Close()
	r.quitM <- true
}

func (r *RabbitMQ) Publish(exchange, routingKey string, body []byte) error {
	return r.Ch.Publish(exchange, routingKey, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        body,
	})
}

func (r *RabbitMQ) Consume(queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args amqp.Table) (<-chan amqp.Delivery, error) {
	return r.Ch.Consume(queue, consumer, autoAck, exclusive, noLocal, noWait, args)
}

func (r *RabbitMQ) QueueDeclare(name string, durable, autoDelete, exclusive, noWait bool, args amqp.Table) (amqp.Queue, error) {
	return r.Ch.QueueDeclare(name, durable, autoDelete, exclusive, noWait, args)
}

func (r *RabbitMQ) ExchangeDeclare(name, kind string, durable, autoDelete, internal, noWait bool, args amqp.Table) error {
	return r.Ch.ExchangeDeclare(name, kind, durable, autoDelete, internal, noWait, args)
}

func (r *RabbitMQ) QueueBind(name, key, exchange string, noWait bool, args amqp.Table) error {
	return r.Ch.QueueBind(name, key, exchange, noWait, args)
}

func (r *RabbitMQ) QueueUnbind(name, key, exchange string, args amqp.Table) error {
	return r.Ch.QueueUnbind(name, key, exchange, args)
}

func (r *RabbitMQ) QueueDelete(name string, ifUnused, ifEmpty, noWait bool) (int, error) {
	return r.Ch.QueueDelete(name, ifUnused, ifEmpty, noWait)
}

func (r *RabbitMQ) ExchangeDelete(name string, ifUnused, noWait bool) error {
	return r.Ch.ExchangeDelete(name, ifUnused, noWait)
}

func (r *RabbitMQ) PublishWithOpts(body []byte) error {
	return r.Ch.Publish(r.Cfg.Exchange, r.Cfg.Key, r.Cfg.Mandatory, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        body,
	})
}

func (r *RabbitMQ) ConsumeWithOpts() (<-chan amqp.Delivery, error) {
	return r.Ch.Consume(r.Cfg.Name, r.Cfg.Consumer,
		r.Cfg.AutoAck, r.Cfg.Exclusive,
		r.Cfg.NoLocal, r.Cfg.NoWait, r.Cfg.Args)
}

func (r *RabbitMQ) QueueDeclareWithOpts() (amqp.Queue, error) {
	q, err := r.Ch.QueueDeclare(r.Cfg.Name, r.Cfg.Durable,
		r.Cfg.AutoDelete, r.Cfg.Exclusive,
		r.Cfg.NoWait, r.Cfg.Args)
	if err != nil {
		return q, err
	}

	if err := r.QueueBind(q.Name, r.Cfg.Key, r.Cfg.Exchange,
		r.Cfg.NoWait, r.Cfg.Args); err != nil {
		return q, err
	}

	return q, nil
}

func (r *RabbitMQ) ExchangeDeclareWithOpts() error {
	return r.Ch.ExchangeDeclare(r.Cfg.Exchange,
		r.Cfg.Kind, r.Cfg.Durable,
		r.Cfg.AutoDelete, r.Cfg.Internal,
		r.Cfg.NoWait, r.Cfg.Args)
}

func (r *RabbitMQ) IsDisconected() bool {
	return r.Conn.IsClosed()
}

func (r *RabbitMQ) ConnectionDown() <-chan bool {
	return r.connDown
}
