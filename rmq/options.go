package rmq

import amqp "github.com/rabbitmq/amqp091-go"

type Options func(*RmqParams) error

type RmqParams struct {
	url        string
	queue      string
	consumer   string
	name       string
	exchange   string
	kind       string
	key        string
	heartbeat  string
	internal   bool
	durable    bool
	autoDelete bool
	exclusive  bool
	noWait     bool
	noLocal    bool
	autoAck    bool
	mandatory  bool
	immediate  bool
	args       amqp.Table
	mock       *MockRabbitMQRepo
}

func newRmqParams(opts ...Options) (*RmqParams, error) {
	r := &RmqParams{}
	for _, opt := range opts {
		if err := opt(r); err != nil {
			return nil, err
		}
	}
	return r, nil

}

func WithMock(mock *MockRabbitMQRepo) Options {
	return func(r *RmqParams) error {
		r.mock = mock
		return nil
	}
}

func WithUrl(url string) Options {
	return func(r *RmqParams) error {
		r.url = url
		return nil
	}
}

func WithQueue(queue string) Options {
	return func(r *RmqParams) error {
		r.queue = queue
		return nil
	}
}

func WithConsumer(consumer string) Options {
	return func(r *RmqParams) error {
		r.consumer = consumer
		return nil
	}
}

func WithName(name string) Options {
	return func(r *RmqParams) error {
		r.name = name
		return nil
	}
}

func WithExchange(exchange string) Options {
	return func(r *RmqParams) error {
		r.exchange = exchange
		return nil
	}
}

func WithKind(kind string) Options {
	return func(r *RmqParams) error {
		r.kind = kind
		return nil
	}
}

func WithKey(key string) Options {
	return func(r *RmqParams) error {
		r.key = key
		return nil
	}
}

func WithHeartbeat(heartbeat string) Options {
	return func(r *RmqParams) error {
		r.heartbeat = heartbeat
		return nil
	}
}

func WithInternal(internal bool) Options {
	return func(r *RmqParams) error {
		r.internal = internal
		return nil
	}
}

func WithDurable(durable bool) Options {
	return func(r *RmqParams) error {
		r.durable = durable
		return nil
	}
}

func WithAutoDelete(autoDelete bool) Options {
	return func(r *RmqParams) error {
		r.autoDelete = autoDelete
		return nil
	}
}

func WithExclusive(exclusive bool) Options {
	return func(r *RmqParams) error {
		r.exclusive = exclusive
		return nil
	}
}

func WithNoWait(noWait bool) Options {
	return func(r *RmqParams) error {
		r.noWait = noWait
		return nil
	}
}

func WithNoLocal(noLocal bool) Options {
	return func(r *RmqParams) error {
		r.noLocal = noLocal
		return nil
	}
}

func WithAutoAck(autoAck bool) Options {
	return func(r *RmqParams) error {
		r.autoAck = autoAck
		return nil
	}
}

func WithMandatory(mandatory bool) Options {
	return func(r *RmqParams) error {
		r.mandatory = mandatory
		return nil
	}
}

func WithImmediate(immediate bool) Options {
	return func(r *RmqParams) error {
		r.immediate = immediate
		return nil
	}
}

func WithArgs(args amqp.Table) Options {
	return func(r *RmqParams) error {
		r.args = args
		return nil
	}
}

func (r *RmqParams) GetMock() *MockRabbitMQRepo {
	return r.mock
}

func (r *RmqParams) GetUrl() string {
	return r.url
}

func (r *RmqParams) GetQueue() string {
	return r.queue
}

func (r *RmqParams) GetConsumer() string {
	return r.consumer
}

func (r *RmqParams) GetName() string {
	return r.name
}

func (r *RmqParams) GetExchange() string {
	return r.exchange
}

func (r *RmqParams) GetKind() string {
	return r.kind
}

func (r *RmqParams) GetKey() string {
	return r.key
}

func (r *RmqParams) GetHeartbeat() string {
	return r.heartbeat
}

func (r *RmqParams) GetInternal() bool {
	return r.internal
}

func (r *RmqParams) GetDurable() bool {
	return r.durable
}

func (r *RmqParams) GetAutoDelete() bool {
	return r.autoDelete
}

func (r *RmqParams) GetExclusive() bool {
	return r.exclusive
}

func (r *RmqParams) GetNoWait() bool {
	return r.noWait
}

func (r *RmqParams) GetNoLocal() bool {
	return r.noLocal
}

func (r *RmqParams) GetAutoAck() bool {
	return r.autoAck
}

func (r *RmqParams) GetMandatory() bool {
	return r.mandatory
}

func (r *RmqParams) GetImmediate() bool {
	return r.immediate
}

func (r *RmqParams) GetArgs() amqp.Table {
	return r.args
}

func (r *RmqParams) GetParams() *RmqParams {
	return r
}

func (r *RmqParams) SetUrl(url string) {
	r.url = url
}

func (r *RmqParams) SetQueue(queue string) {
	r.queue = queue
}

func (r *RmqParams) SetConsumer(consumer string) {
	r.consumer = consumer
}

func (r *RmqParams) SetName(name string) {
	r.name = name
}

func (r *RmqParams) SetExchange(exchange string) {
	r.exchange = exchange
}

func (r *RmqParams) SetKind(kind string) {
	r.kind = kind
}

func (r *RmqParams) SetKey(key string) {
	r.key = key
}

func (r *RmqParams) SetHeartbeat(heartbeat string) {
	r.heartbeat = heartbeat
}

func (r *RmqParams) SetInternal(internal bool) {
	r.internal = internal
}

func (r *RmqParams) SetDurable(durable bool) {
	r.durable = durable
}

func (r *RmqParams) SetAutoDelete(autoDelete bool) {
	r.autoDelete = autoDelete
}

func (r *RmqParams) SetExclusive(exclusive bool) {
	r.exclusive = exclusive
}

func (r *RmqParams) SetNoWait(noWait bool) {
	r.noWait = noWait
}

func (r *RmqParams) SetNoLocal(noLocal bool) {
	r.noLocal = noLocal
}

func (r *RmqParams) SetAutoAck(autoAck bool) {
	r.autoAck = autoAck
}

func (r *RmqParams) SetMandatory(mandatory bool) {
	r.mandatory = mandatory
}

func (r *RmqParams) SetImmediate(immediate bool) {
	r.immediate = immediate
}

func (r *RmqParams) SetArgs(args amqp.Table) {
	r.args = args
}
