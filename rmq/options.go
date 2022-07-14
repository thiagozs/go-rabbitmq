package rmq

import amqp "github.com/rabbitmq/amqp091-go"

type Options func(*Config)

type Config struct {
	Url        string
	Queue      string
	Consumer   string
	Name       string
	Exchange   string
	Kind       string
	Key        string
	Heartbeat  string
	Internal   bool
	Durable    bool
	AutoDelete bool
	Exclusive  bool
	NoWait     bool
	NoLocal    bool
	AutoAck    bool
	Args       amqp.Table
}

func NewOptions(opts ...Options) Config {
	opt := Config{}
	for _, o := range opts {
		o(&opt)
	}
	return opt
}

func Name(name string) Options {
	return func(o *Config) {
		o.Name = name
	}
}

func Durable(t bool) Options {
	return func(o *Config) {
		o.Durable = t
	}
}

func AutoDelete(t bool) Options {
	return func(o *Config) {
		o.AutoDelete = t
	}
}

func Exclusive(t bool) Options {
	return func(o *Config) {
		o.Exclusive = t
	}
}

func NoWait(t bool) Options {
	return func(o *Config) {
		o.NoWait = t
	}
}

func NoLocal(t bool) Options {
	return func(o *Config) {
		o.NoLocal = t
	}
}

func AutoAck(t bool) Options {
	return func(o *Config) {
		o.AutoAck = t
	}
}

func Url(url string) Options {
	return func(o *Config) {
		o.Url = url
	}
}

func Queue(name string) Options {
	return func(o *Config) {
		o.Name = name
	}
}

func Consumer(name string) Options {
	return func(o *Config) {
		o.Consumer = name
	}
}

func Exchange(name string) Options {
	return func(o *Config) {
		o.Exchange = name
	}
}

func Kind(kind string) Options {
	return func(o *Config) {
		o.Kind = kind
	}
}

func Internal(t bool) Options {
	return func(o *Config) {
		o.Internal = t
	}
}

func Args(args map[string]interface{}) Options {
	return func(o *Config) {
		o.Args = args
	}
}

func Key(key string) Options {
	return func(o *Config) {
		o.Key = key
	}
}

func Heartbeat(heartbeat string) Options {
	return func(o *Config) {
		o.Heartbeat = heartbeat
	}
}
