// Code generated by MockGen. DO NOT EDIT.
// Source: chan_repo.go
//
// Generated by this command:
//
//	mockgen -source=chan_repo.go -package=rmq -destination=chan_mock.go
//

// Package rmq is a generated GoMock package.
package rmq

import (
	reflect "reflect"

	amqp091 "github.com/rabbitmq/amqp091-go"
	gomock "go.uber.org/mock/gomock"
)

// MockChannelRepo is a mock of ChannelRepo interface.
type MockChannelRepo struct {
	ctrl     *gomock.Controller
	recorder *MockChannelRepoMockRecorder
}

// MockChannelRepoMockRecorder is the mock recorder for MockChannelRepo.
type MockChannelRepoMockRecorder struct {
	mock *MockChannelRepo
}

// NewMockChannelRepo creates a new mock instance.
func NewMockChannelRepo(ctrl *gomock.Controller) *MockChannelRepo {
	mock := &MockChannelRepo{ctrl: ctrl}
	mock.recorder = &MockChannelRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChannelRepo) EXPECT() *MockChannelRepoMockRecorder {
	return m.recorder
}

// Ack mocks base method.
func (m *MockChannelRepo) Ack(tag uint64, multiple bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ack", tag, multiple)
	ret0, _ := ret[0].(error)
	return ret0
}

// Ack indicates an expected call of Ack.
func (mr *MockChannelRepoMockRecorder) Ack(tag, multiple any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ack", reflect.TypeOf((*MockChannelRepo)(nil).Ack), tag, multiple)
}

// Cancel mocks base method.
func (m *MockChannelRepo) Cancel(consumer string, noWait bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cancel", consumer, noWait)
	ret0, _ := ret[0].(error)
	return ret0
}

// Cancel indicates an expected call of Cancel.
func (mr *MockChannelRepoMockRecorder) Cancel(consumer, noWait any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cancel", reflect.TypeOf((*MockChannelRepo)(nil).Cancel), consumer, noWait)
}

// Close mocks base method.
func (m *MockChannelRepo) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockChannelRepoMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockChannelRepo)(nil).Close))
}

// Confirm mocks base method.
func (m *MockChannelRepo) Confirm(noWait bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Confirm", noWait)
	ret0, _ := ret[0].(error)
	return ret0
}

// Confirm indicates an expected call of Confirm.
func (mr *MockChannelRepoMockRecorder) Confirm(noWait any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Confirm", reflect.TypeOf((*MockChannelRepo)(nil).Confirm), noWait)
}

// Consume mocks base method.
func (m *MockChannelRepo) Consume(queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args amqp091.Table) (<-chan amqp091.Delivery, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Consume", queue, consumer, autoAck, exclusive, noLocal, noWait, args)
	ret0, _ := ret[0].(<-chan amqp091.Delivery)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Consume indicates an expected call of Consume.
func (mr *MockChannelRepoMockRecorder) Consume(queue, consumer, autoAck, exclusive, noLocal, noWait, args any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Consume", reflect.TypeOf((*MockChannelRepo)(nil).Consume), queue, consumer, autoAck, exclusive, noLocal, noWait, args)
}

// ExchangeBind mocks base method.
func (m *MockChannelRepo) ExchangeBind(destination, key, source string, noWait bool, args amqp091.Table) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExchangeBind", destination, key, source, noWait, args)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExchangeBind indicates an expected call of ExchangeBind.
func (mr *MockChannelRepoMockRecorder) ExchangeBind(destination, key, source, noWait, args any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExchangeBind", reflect.TypeOf((*MockChannelRepo)(nil).ExchangeBind), destination, key, source, noWait, args)
}

// ExchangeDeclare mocks base method.
func (m *MockChannelRepo) ExchangeDeclare(name, kind string, durable, autoDelete, internal, noWait bool, args amqp091.Table) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExchangeDeclare", name, kind, durable, autoDelete, internal, noWait, args)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExchangeDeclare indicates an expected call of ExchangeDeclare.
func (mr *MockChannelRepoMockRecorder) ExchangeDeclare(name, kind, durable, autoDelete, internal, noWait, args any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExchangeDeclare", reflect.TypeOf((*MockChannelRepo)(nil).ExchangeDeclare), name, kind, durable, autoDelete, internal, noWait, args)
}

// ExchangeDeclarePassive mocks base method.
func (m *MockChannelRepo) ExchangeDeclarePassive(name, kind string, durable, autoDelete, internal, noWait bool, args amqp091.Table) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExchangeDeclarePassive", name, kind, durable, autoDelete, internal, noWait, args)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExchangeDeclarePassive indicates an expected call of ExchangeDeclarePassive.
func (mr *MockChannelRepoMockRecorder) ExchangeDeclarePassive(name, kind, durable, autoDelete, internal, noWait, args any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExchangeDeclarePassive", reflect.TypeOf((*MockChannelRepo)(nil).ExchangeDeclarePassive), name, kind, durable, autoDelete, internal, noWait, args)
}

// ExchangeDelete mocks base method.
func (m *MockChannelRepo) ExchangeDelete(name string, ifUnused, noWait bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExchangeDelete", name, ifUnused, noWait)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExchangeDelete indicates an expected call of ExchangeDelete.
func (mr *MockChannelRepoMockRecorder) ExchangeDelete(name, ifUnused, noWait any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExchangeDelete", reflect.TypeOf((*MockChannelRepo)(nil).ExchangeDelete), name, ifUnused, noWait)
}

// ExchangeUnbind mocks base method.
func (m *MockChannelRepo) ExchangeUnbind(destination, key, source string, noWait bool, args amqp091.Table) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExchangeUnbind", destination, key, source, noWait, args)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExchangeUnbind indicates an expected call of ExchangeUnbind.
func (mr *MockChannelRepoMockRecorder) ExchangeUnbind(destination, key, source, noWait, args any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExchangeUnbind", reflect.TypeOf((*MockChannelRepo)(nil).ExchangeUnbind), destination, key, source, noWait, args)
}

// Flow mocks base method.
func (m *MockChannelRepo) Flow(active bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Flow", active)
	ret0, _ := ret[0].(error)
	return ret0
}

// Flow indicates an expected call of Flow.
func (mr *MockChannelRepoMockRecorder) Flow(active any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Flow", reflect.TypeOf((*MockChannelRepo)(nil).Flow), active)
}

// Get mocks base method.
func (m *MockChannelRepo) Get(queue string, autoAck bool) (amqp091.Delivery, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", queue, autoAck)
	ret0, _ := ret[0].(amqp091.Delivery)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get.
func (mr *MockChannelRepoMockRecorder) Get(queue, autoAck any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockChannelRepo)(nil).Get), queue, autoAck)
}

// GetNextPublishSeqNo mocks base method.
func (m *MockChannelRepo) GetNextPublishSeqNo() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNextPublishSeqNo")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetNextPublishSeqNo indicates an expected call of GetNextPublishSeqNo.
func (mr *MockChannelRepoMockRecorder) GetNextPublishSeqNo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNextPublishSeqNo", reflect.TypeOf((*MockChannelRepo)(nil).GetNextPublishSeqNo))
}

// IsClosed mocks base method.
func (m *MockChannelRepo) IsClosed() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsClosed")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsClosed indicates an expected call of IsClosed.
func (mr *MockChannelRepoMockRecorder) IsClosed() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsClosed", reflect.TypeOf((*MockChannelRepo)(nil).IsClosed))
}

// Nack mocks base method.
func (m *MockChannelRepo) Nack(tag uint64, multiple, requeue bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Nack", tag, multiple, requeue)
	ret0, _ := ret[0].(error)
	return ret0
}

// Nack indicates an expected call of Nack.
func (mr *MockChannelRepoMockRecorder) Nack(tag, multiple, requeue any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Nack", reflect.TypeOf((*MockChannelRepo)(nil).Nack), tag, multiple, requeue)
}

// NotifyCancel mocks base method.
func (m *MockChannelRepo) NotifyCancel(c chan string) chan string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NotifyCancel", c)
	ret0, _ := ret[0].(chan string)
	return ret0
}

// NotifyCancel indicates an expected call of NotifyCancel.
func (mr *MockChannelRepoMockRecorder) NotifyCancel(c any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyCancel", reflect.TypeOf((*MockChannelRepo)(nil).NotifyCancel), c)
}

// NotifyClose mocks base method.
func (m *MockChannelRepo) NotifyClose(c chan *amqp091.Error) chan *amqp091.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NotifyClose", c)
	ret0, _ := ret[0].(chan *amqp091.Error)
	return ret0
}

// NotifyClose indicates an expected call of NotifyClose.
func (mr *MockChannelRepoMockRecorder) NotifyClose(c any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyClose", reflect.TypeOf((*MockChannelRepo)(nil).NotifyClose), c)
}

// NotifyConfirm mocks base method.
func (m *MockChannelRepo) NotifyConfirm(ack, nack chan uint64) (chan uint64, chan uint64) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NotifyConfirm", ack, nack)
	ret0, _ := ret[0].(chan uint64)
	ret1, _ := ret[1].(chan uint64)
	return ret0, ret1
}

// NotifyConfirm indicates an expected call of NotifyConfirm.
func (mr *MockChannelRepoMockRecorder) NotifyConfirm(ack, nack any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyConfirm", reflect.TypeOf((*MockChannelRepo)(nil).NotifyConfirm), ack, nack)
}

// NotifyFlow mocks base method.
func (m *MockChannelRepo) NotifyFlow(c chan bool) chan bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NotifyFlow", c)
	ret0, _ := ret[0].(chan bool)
	return ret0
}

// NotifyFlow indicates an expected call of NotifyFlow.
func (mr *MockChannelRepoMockRecorder) NotifyFlow(c any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyFlow", reflect.TypeOf((*MockChannelRepo)(nil).NotifyFlow), c)
}

// NotifyPublish mocks base method.
func (m *MockChannelRepo) NotifyPublish(confirm chan amqp091.Confirmation) chan amqp091.Confirmation {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NotifyPublish", confirm)
	ret0, _ := ret[0].(chan amqp091.Confirmation)
	return ret0
}

// NotifyPublish indicates an expected call of NotifyPublish.
func (mr *MockChannelRepoMockRecorder) NotifyPublish(confirm any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyPublish", reflect.TypeOf((*MockChannelRepo)(nil).NotifyPublish), confirm)
}

// NotifyReturn mocks base method.
func (m *MockChannelRepo) NotifyReturn(c chan amqp091.Return) chan amqp091.Return {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NotifyReturn", c)
	ret0, _ := ret[0].(chan amqp091.Return)
	return ret0
}

// NotifyReturn indicates an expected call of NotifyReturn.
func (mr *MockChannelRepoMockRecorder) NotifyReturn(c any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyReturn", reflect.TypeOf((*MockChannelRepo)(nil).NotifyReturn), c)
}

// Publish mocks base method.
func (m *MockChannelRepo) Publish(exchange, key string, mandatory, immediate bool, msg amqp091.Publishing) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publish", exchange, key, mandatory, immediate, msg)
	ret0, _ := ret[0].(error)
	return ret0
}

// Publish indicates an expected call of Publish.
func (mr *MockChannelRepoMockRecorder) Publish(exchange, key, mandatory, immediate, msg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockChannelRepo)(nil).Publish), exchange, key, mandatory, immediate, msg)
}

// PublishWithDeferredConfirm mocks base method.
func (m *MockChannelRepo) PublishWithDeferredConfirm(exchange, key string, mandatory, immediate bool, msg amqp091.Publishing) (*amqp091.DeferredConfirmation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishWithDeferredConfirm", exchange, key, mandatory, immediate, msg)
	ret0, _ := ret[0].(*amqp091.DeferredConfirmation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PublishWithDeferredConfirm indicates an expected call of PublishWithDeferredConfirm.
func (mr *MockChannelRepoMockRecorder) PublishWithDeferredConfirm(exchange, key, mandatory, immediate, msg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishWithDeferredConfirm", reflect.TypeOf((*MockChannelRepo)(nil).PublishWithDeferredConfirm), exchange, key, mandatory, immediate, msg)
}

// Qos mocks base method.
func (m *MockChannelRepo) Qos(prefetchCount, prefetchSize int, global bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Qos", prefetchCount, prefetchSize, global)
	ret0, _ := ret[0].(error)
	return ret0
}

// Qos indicates an expected call of Qos.
func (mr *MockChannelRepoMockRecorder) Qos(prefetchCount, prefetchSize, global any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Qos", reflect.TypeOf((*MockChannelRepo)(nil).Qos), prefetchCount, prefetchSize, global)
}

// QueueBind mocks base method.
func (m *MockChannelRepo) QueueBind(name, key, exchange string, noWait bool, args amqp091.Table) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueueBind", name, key, exchange, noWait, args)
	ret0, _ := ret[0].(error)
	return ret0
}

// QueueBind indicates an expected call of QueueBind.
func (mr *MockChannelRepoMockRecorder) QueueBind(name, key, exchange, noWait, args any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueueBind", reflect.TypeOf((*MockChannelRepo)(nil).QueueBind), name, key, exchange, noWait, args)
}

// QueueDeclare mocks base method.
func (m *MockChannelRepo) QueueDeclare(name string, durable, autoDelete, exclusive, noWait bool, args amqp091.Table) (amqp091.Queue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueueDeclare", name, durable, autoDelete, exclusive, noWait, args)
	ret0, _ := ret[0].(amqp091.Queue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueueDeclare indicates an expected call of QueueDeclare.
func (mr *MockChannelRepoMockRecorder) QueueDeclare(name, durable, autoDelete, exclusive, noWait, args any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueueDeclare", reflect.TypeOf((*MockChannelRepo)(nil).QueueDeclare), name, durable, autoDelete, exclusive, noWait, args)
}

// QueueDeclarePassive mocks base method.
func (m *MockChannelRepo) QueueDeclarePassive(name string, durable, autoDelete, exclusive, noWait bool, args amqp091.Table) (amqp091.Queue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueueDeclarePassive", name, durable, autoDelete, exclusive, noWait, args)
	ret0, _ := ret[0].(amqp091.Queue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueueDeclarePassive indicates an expected call of QueueDeclarePassive.
func (mr *MockChannelRepoMockRecorder) QueueDeclarePassive(name, durable, autoDelete, exclusive, noWait, args any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueueDeclarePassive", reflect.TypeOf((*MockChannelRepo)(nil).QueueDeclarePassive), name, durable, autoDelete, exclusive, noWait, args)
}

// QueueDelete mocks base method.
func (m *MockChannelRepo) QueueDelete(name string, ifUnused, ifEmpty, noWait bool) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueueDelete", name, ifUnused, ifEmpty, noWait)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueueDelete indicates an expected call of QueueDelete.
func (mr *MockChannelRepoMockRecorder) QueueDelete(name, ifUnused, ifEmpty, noWait any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueueDelete", reflect.TypeOf((*MockChannelRepo)(nil).QueueDelete), name, ifUnused, ifEmpty, noWait)
}

// QueueInspect mocks base method.
func (m *MockChannelRepo) QueueInspect(name string) (amqp091.Queue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueueInspect", name)
	ret0, _ := ret[0].(amqp091.Queue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueueInspect indicates an expected call of QueueInspect.
func (mr *MockChannelRepoMockRecorder) QueueInspect(name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueueInspect", reflect.TypeOf((*MockChannelRepo)(nil).QueueInspect), name)
}

// QueuePurge mocks base method.
func (m *MockChannelRepo) QueuePurge(name string, noWait bool) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueuePurge", name, noWait)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueuePurge indicates an expected call of QueuePurge.
func (mr *MockChannelRepoMockRecorder) QueuePurge(name, noWait any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueuePurge", reflect.TypeOf((*MockChannelRepo)(nil).QueuePurge), name, noWait)
}

// QueueUnbind mocks base method.
func (m *MockChannelRepo) QueueUnbind(name, key, exchange string, args amqp091.Table) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueueUnbind", name, key, exchange, args)
	ret0, _ := ret[0].(error)
	return ret0
}

// QueueUnbind indicates an expected call of QueueUnbind.
func (mr *MockChannelRepoMockRecorder) QueueUnbind(name, key, exchange, args any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueueUnbind", reflect.TypeOf((*MockChannelRepo)(nil).QueueUnbind), name, key, exchange, args)
}

// Recover mocks base method.
func (m *MockChannelRepo) Recover(requeue bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recover", requeue)
	ret0, _ := ret[0].(error)
	return ret0
}

// Recover indicates an expected call of Recover.
func (mr *MockChannelRepoMockRecorder) Recover(requeue any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recover", reflect.TypeOf((*MockChannelRepo)(nil).Recover), requeue)
}

// Reject mocks base method.
func (m *MockChannelRepo) Reject(tag uint64, requeue bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reject", tag, requeue)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reject indicates an expected call of Reject.
func (mr *MockChannelRepoMockRecorder) Reject(tag, requeue any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reject", reflect.TypeOf((*MockChannelRepo)(nil).Reject), tag, requeue)
}

// Tx mocks base method.
func (m *MockChannelRepo) Tx() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tx")
	ret0, _ := ret[0].(error)
	return ret0
}

// Tx indicates an expected call of Tx.
func (mr *MockChannelRepoMockRecorder) Tx() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tx", reflect.TypeOf((*MockChannelRepo)(nil).Tx))
}

// TxCommit mocks base method.
func (m *MockChannelRepo) TxCommit() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TxCommit")
	ret0, _ := ret[0].(error)
	return ret0
}

// TxCommit indicates an expected call of TxCommit.
func (mr *MockChannelRepoMockRecorder) TxCommit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TxCommit", reflect.TypeOf((*MockChannelRepo)(nil).TxCommit))
}

// TxRollback mocks base method.
func (m *MockChannelRepo) TxRollback() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TxRollback")
	ret0, _ := ret[0].(error)
	return ret0
}

// TxRollback indicates an expected call of TxRollback.
func (mr *MockChannelRepoMockRecorder) TxRollback() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TxRollback", reflect.TypeOf((*MockChannelRepo)(nil).TxRollback))
}