// Code generated by MockGen. DO NOT EDIT.
// Source: rmq_repo.go
//
// Generated by this command:
//
//	mockgen -source=rmq_repo.go -package=rmq -destination=rmq_mock.go
//

// Package rmq is a generated GoMock package.
package rmq

import (
	reflect "reflect"

	amqp091 "github.com/rabbitmq/amqp091-go"
	gomock "go.uber.org/mock/gomock"
)

// MockRabbitMQRepo is a mock of RabbitMQRepo interface.
type MockRabbitMQRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRabbitMQRepoMockRecorder
}

// MockRabbitMQRepoMockRecorder is the mock recorder for MockRabbitMQRepo.
type MockRabbitMQRepoMockRecorder struct {
	mock *MockRabbitMQRepo
}

// NewMockRabbitMQRepo creates a new mock instance.
func NewMockRabbitMQRepo(ctrl *gomock.Controller) *MockRabbitMQRepo {
	mock := &MockRabbitMQRepo{ctrl: ctrl}
	mock.recorder = &MockRabbitMQRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRabbitMQRepo) EXPECT() *MockRabbitMQRepoMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockRabbitMQRepo) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockRabbitMQRepoMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockRabbitMQRepo)(nil).Close))
}

// Consume mocks base method.
func (m *MockRabbitMQRepo) Consume(queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args amqp091.Table) (<-chan amqp091.Delivery, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Consume", queue, consumer, autoAck, exclusive, noLocal, noWait, args)
	ret0, _ := ret[0].(<-chan amqp091.Delivery)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Consume indicates an expected call of Consume.
func (mr *MockRabbitMQRepoMockRecorder) Consume(queue, consumer, autoAck, exclusive, noLocal, noWait, args any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Consume", reflect.TypeOf((*MockRabbitMQRepo)(nil).Consume), queue, consumer, autoAck, exclusive, noLocal, noWait, args)
}

// ConsumeWithOpts mocks base method.
func (m *MockRabbitMQRepo) ConsumeWithOpts() (<-chan amqp091.Delivery, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConsumeWithOpts")
	ret0, _ := ret[0].(<-chan amqp091.Delivery)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConsumeWithOpts indicates an expected call of ConsumeWithOpts.
func (mr *MockRabbitMQRepoMockRecorder) ConsumeWithOpts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConsumeWithOpts", reflect.TypeOf((*MockRabbitMQRepo)(nil).ConsumeWithOpts))
}

// ExchangeDeclare mocks base method.
func (m *MockRabbitMQRepo) ExchangeDeclare(name, kind string, durable, autoDelete, internal, noWait bool, args amqp091.Table) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExchangeDeclare", name, kind, durable, autoDelete, internal, noWait, args)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExchangeDeclare indicates an expected call of ExchangeDeclare.
func (mr *MockRabbitMQRepoMockRecorder) ExchangeDeclare(name, kind, durable, autoDelete, internal, noWait, args any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExchangeDeclare", reflect.TypeOf((*MockRabbitMQRepo)(nil).ExchangeDeclare), name, kind, durable, autoDelete, internal, noWait, args)
}

// ExchangeDeclareWithOpts mocks base method.
func (m *MockRabbitMQRepo) ExchangeDeclareWithOpts() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExchangeDeclareWithOpts")
	ret0, _ := ret[0].(error)
	return ret0
}

// ExchangeDeclareWithOpts indicates an expected call of ExchangeDeclareWithOpts.
func (mr *MockRabbitMQRepoMockRecorder) ExchangeDeclareWithOpts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExchangeDeclareWithOpts", reflect.TypeOf((*MockRabbitMQRepo)(nil).ExchangeDeclareWithOpts))
}

// ExchangeDelete mocks base method.
func (m *MockRabbitMQRepo) ExchangeDelete(name string, ifUnused, noWait bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExchangeDelete", name, ifUnused, noWait)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExchangeDelete indicates an expected call of ExchangeDelete.
func (mr *MockRabbitMQRepoMockRecorder) ExchangeDelete(name, ifUnused, noWait any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExchangeDelete", reflect.TypeOf((*MockRabbitMQRepo)(nil).ExchangeDelete), name, ifUnused, noWait)
}

// GetCh mocks base method.
func (m *MockRabbitMQRepo) GetCh() ChannelRepo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCh")
	ret0, _ := ret[0].(ChannelRepo)
	return ret0
}

// GetCh indicates an expected call of GetCh.
func (mr *MockRabbitMQRepoMockRecorder) GetCh() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCh", reflect.TypeOf((*MockRabbitMQRepo)(nil).GetCh))
}

// GetConn mocks base method.
func (m *MockRabbitMQRepo) GetConn() ConnectionRepo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConn")
	ret0, _ := ret[0].(ConnectionRepo)
	return ret0
}

// GetConn indicates an expected call of GetConn.
func (mr *MockRabbitMQRepoMockRecorder) GetConn() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConn", reflect.TypeOf((*MockRabbitMQRepo)(nil).GetConn))
}

// IsDisconected mocks base method.
func (m *MockRabbitMQRepo) IsDisconected() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsDisconected")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsDisconected indicates an expected call of IsDisconected.
func (mr *MockRabbitMQRepoMockRecorder) IsDisconected() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsDisconected", reflect.TypeOf((*MockRabbitMQRepo)(nil).IsDisconected))
}

// MonitorConn mocks base method.
func (m *MockRabbitMQRepo) MonitorConn() <-chan bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MonitorConn")
	ret0, _ := ret[0].(<-chan bool)
	return ret0
}

// MonitorConn indicates an expected call of MonitorConn.
func (mr *MockRabbitMQRepoMockRecorder) MonitorConn() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MonitorConn", reflect.TypeOf((*MockRabbitMQRepo)(nil).MonitorConn))
}

// Publish mocks base method.
func (m *MockRabbitMQRepo) Publish(exchange, routingKey string, body []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publish", exchange, routingKey, body)
	ret0, _ := ret[0].(error)
	return ret0
}

// Publish indicates an expected call of Publish.
func (mr *MockRabbitMQRepoMockRecorder) Publish(exchange, routingKey, body any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockRabbitMQRepo)(nil).Publish), exchange, routingKey, body)
}

// PublishWithOpts mocks base method.
func (m *MockRabbitMQRepo) PublishWithOpts(body []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishWithOpts", body)
	ret0, _ := ret[0].(error)
	return ret0
}

// PublishWithOpts indicates an expected call of PublishWithOpts.
func (mr *MockRabbitMQRepoMockRecorder) PublishWithOpts(body any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishWithOpts", reflect.TypeOf((*MockRabbitMQRepo)(nil).PublishWithOpts), body)
}

// QueueBind mocks base method.
func (m *MockRabbitMQRepo) QueueBind(name, key, exchange string, noWait bool, args amqp091.Table) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueueBind", name, key, exchange, noWait, args)
	ret0, _ := ret[0].(error)
	return ret0
}

// QueueBind indicates an expected call of QueueBind.
func (mr *MockRabbitMQRepoMockRecorder) QueueBind(name, key, exchange, noWait, args any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueueBind", reflect.TypeOf((*MockRabbitMQRepo)(nil).QueueBind), name, key, exchange, noWait, args)
}

// QueueDeclare mocks base method.
func (m *MockRabbitMQRepo) QueueDeclare(name string, durable, autoDelete, exclusive, noWait bool, args amqp091.Table) (amqp091.Queue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueueDeclare", name, durable, autoDelete, exclusive, noWait, args)
	ret0, _ := ret[0].(amqp091.Queue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueueDeclare indicates an expected call of QueueDeclare.
func (mr *MockRabbitMQRepoMockRecorder) QueueDeclare(name, durable, autoDelete, exclusive, noWait, args any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueueDeclare", reflect.TypeOf((*MockRabbitMQRepo)(nil).QueueDeclare), name, durable, autoDelete, exclusive, noWait, args)
}

// QueueDeclareWithOpts mocks base method.
func (m *MockRabbitMQRepo) QueueDeclareWithOpts() (amqp091.Queue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueueDeclareWithOpts")
	ret0, _ := ret[0].(amqp091.Queue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueueDeclareWithOpts indicates an expected call of QueueDeclareWithOpts.
func (mr *MockRabbitMQRepoMockRecorder) QueueDeclareWithOpts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueueDeclareWithOpts", reflect.TypeOf((*MockRabbitMQRepo)(nil).QueueDeclareWithOpts))
}

// QueueDelete mocks base method.
func (m *MockRabbitMQRepo) QueueDelete(name string, ifUnused, ifEmpty, noWait bool) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueueDelete", name, ifUnused, ifEmpty, noWait)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueueDelete indicates an expected call of QueueDelete.
func (mr *MockRabbitMQRepoMockRecorder) QueueDelete(name, ifUnused, ifEmpty, noWait any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueueDelete", reflect.TypeOf((*MockRabbitMQRepo)(nil).QueueDelete), name, ifUnused, ifEmpty, noWait)
}

// QueueUnbind mocks base method.
func (m *MockRabbitMQRepo) QueueUnbind(name, key, exchange string, args amqp091.Table) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueueUnbind", name, key, exchange, args)
	ret0, _ := ret[0].(error)
	return ret0
}

// QueueUnbind indicates an expected call of QueueUnbind.
func (mr *MockRabbitMQRepoMockRecorder) QueueUnbind(name, key, exchange, args any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueueUnbind", reflect.TypeOf((*MockRabbitMQRepo)(nil).QueueUnbind), name, key, exchange, args)
}
