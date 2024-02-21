package rmq

import (
	"crypto/tls"
	"net"

	amqp "github.com/rabbitmq/amqp091-go"
)

type ConnectionRepo interface {
	Channel() (*amqp.Channel, error)
	Close() error
	ConnectionState() tls.ConnectionState
	IsClosed() bool
	LocalAddr() net.Addr
	NotifyBlocked(receiver chan amqp.Blocking) chan amqp.Blocking
	NotifyClose(receiver chan *amqp.Error) chan *amqp.Error
}
