package rmq

import (
	"testing"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestNewService(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := NewMockRabbitMQRepo(ctrl)

	mock.EXPECT().GetCh().Return(nil).AnyTimes()
	mock.EXPECT().GetConn().Return(nil).AnyTimes()

	service, err := NewService(WithMock(mock))

	assert.NoError(t, err)
	assert.NotNil(t, service)
}

func TestPublishWithOpts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCh := NewMockChannelRepo(ctrl)
	mockConn := NewMockConnectionRepo(ctrl)
	mockRabbitMQ := NewMockRabbitMQRepo(ctrl)

	mockCh.EXPECT().Publish(gomock.Any(), gomock.Any(),
		gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
	mockRabbitMQ.EXPECT().GetCh().Return(mockCh).AnyTimes()
	mockRabbitMQ.EXPECT().GetConn().Return(mockConn).AnyTimes()

	opts := []Options{
		WithMock(mockRabbitMQ),
	}

	service, err := NewService(opts...)
	assert.NoError(t, err)
	assert.NotNil(t, service)

	err = service.PublishWithOpts([]byte("Hello World"))
	assert.NoError(t, err)
}

func TestConsumeWithOpts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCh := NewMockChannelRepo(ctrl)
	mockConn := NewMockConnectionRepo(ctrl)
	mockRabbitMQ := NewMockRabbitMQRepo(ctrl)

	mockCh.EXPECT().Consume(gomock.Any(), gomock.Any(),
		gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, nil).AnyTimes()
	mockRabbitMQ.EXPECT().GetCh().Return(mockCh).AnyTimes()
	mockRabbitMQ.EXPECT().GetConn().Return(mockConn).AnyTimes()

	opts := []Options{
		WithMock(mockRabbitMQ),
	}

	service, err := NewService(opts...)
	assert.NoError(t, err)
	assert.NotNil(t, service)

	_, err = service.ConsumeWithOpts()
	assert.NoError(t, err)
}

func TestQueueDeclareWithOpts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCh := NewMockChannelRepo(ctrl)
	mockConn := NewMockConnectionRepo(ctrl)
	mockRabbitMQ := NewMockRabbitMQRepo(ctrl)

	mockCh.EXPECT().QueueDeclare(gomock.Any(), gomock.Any(),
		gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(amqp.Queue{}, nil).AnyTimes()
	mockCh.EXPECT().QueueBind(gomock.Any(), gomock.Any(),
		gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
	mockRabbitMQ.EXPECT().GetCh().Return(mockCh).AnyTimes()
	mockRabbitMQ.EXPECT().GetConn().Return(mockConn).AnyTimes()

	opts := []Options{
		WithMock(mockRabbitMQ),
	}

	service, err := NewService(opts...)
	assert.NoError(t, err)
	assert.NotNil(t, service)

	_, err = service.QueueDeclareWithOpts()
	assert.NoError(t, err)
}

func TestExchangeDeclareWithOpts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCh := NewMockChannelRepo(ctrl)
	mockConn := NewMockConnectionRepo(ctrl)
	mockRabbitMQ := NewMockRabbitMQRepo(ctrl)

	mockCh.EXPECT().ExchangeDeclare(gomock.Any(), gomock.Any(),
		gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
	mockRabbitMQ.EXPECT().GetCh().Return(mockCh).AnyTimes()
	mockRabbitMQ.EXPECT().GetConn().Return(mockConn).AnyTimes()

	opts := []Options{
		WithMock(mockRabbitMQ),
	}

	service, err := NewService(opts...)
	assert.NoError(t, err)
	assert.NotNil(t, service)

	err = service.ExchangeDeclareWithOpts()
	assert.NoError(t, err)
}

func TestIsDisconected(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConn := NewMockConnectionRepo(ctrl)
	mockRabbitMQ := NewMockRabbitMQRepo(ctrl)

	mockConn.EXPECT().IsClosed().Return(false).AnyTimes()
	mockRabbitMQ.EXPECT().GetCh().Return(nil).AnyTimes()
	mockRabbitMQ.EXPECT().GetConn().Return(mockConn).AnyTimes()

	opts := []Options{
		WithMock(mockRabbitMQ),
	}

	service, err := NewService(opts...)
	assert.NoError(t, err)
	assert.NotNil(t, service)

	assert.False(t, service.IsDisconected())
}

func TestClose(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCh := NewMockChannelRepo(ctrl)
	mockConn := NewMockConnectionRepo(ctrl)
	mockRabbitMQ := NewMockRabbitMQRepo(ctrl)

	mockCh.EXPECT().Close().Return(nil).AnyTimes()
	mockConn.EXPECT().Close().Return(nil).AnyTimes()

	mockRabbitMQ.EXPECT().GetCh().Return(mockCh).AnyTimes()
	mockRabbitMQ.EXPECT().GetConn().Return(mockConn).AnyTimes()

	mockRabbitMQ.EXPECT().Close().Return(nil).AnyTimes()

	opts := []Options{
		WithMock(mockRabbitMQ),
	}

	service, err := NewService(opts...)
	assert.NoError(t, err)
	assert.NotNil(t, service)

	assert.NoError(t, service.Close())
}

func TestGetConnection(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCh := NewMockChannelRepo(ctrl)
	mockConn := NewMockConnectionRepo(ctrl)
	mockRabbitMQ := NewMockRabbitMQRepo(ctrl)

	mockConn.EXPECT().Close().Return(nil).AnyTimes()
	mockRabbitMQ.EXPECT().GetCh().Return(mockCh).AnyTimes()
	mockRabbitMQ.EXPECT().GetConn().Return(mockConn).AnyTimes()

	opts := []Options{
		WithMock(mockRabbitMQ),
	}

	service, err := NewService(opts...)
	assert.NoError(t, err)
	assert.NotNil(t, service)

	conn, ch, err := service.getConnection()
	assert.NoError(t, err)
	assert.NotNil(t, conn)
	assert.NotNil(t, ch)
}

func TestMonitoringConn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCh := NewMockChannelRepo(ctrl)
	mockConn := NewMockConnectionRepo(ctrl)
	mockRabbitMQ := NewMockRabbitMQRepo(ctrl)

	mockConn.EXPECT().Close().Return(nil).AnyTimes()
	mockRabbitMQ.EXPECT().GetCh().Return(mockCh).AnyTimes()
	mockRabbitMQ.EXPECT().GetConn().Return(mockConn).AnyTimes()

	opts := []Options{
		WithMock(mockRabbitMQ),
	}

	service, err := NewService(opts...)
	assert.NoError(t, err)
	assert.NotNil(t, service)

	go service.monitoringConn(service.getConnection)
}

func TestMonitorConn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCh := NewMockChannelRepo(ctrl)
	mockConn := NewMockConnectionRepo(ctrl)
	mockRabbitMQ := NewMockRabbitMQRepo(ctrl)

	mockConn.EXPECT().Close().Return(nil).AnyTimes()
	mockRabbitMQ.EXPECT().GetCh().Return(mockCh).AnyTimes()
	mockRabbitMQ.EXPECT().GetConn().Return(mockConn).AnyTimes()

	mockRabbitMQ.EXPECT().MonitorConn().Return(make(chan bool)).AnyTimes()

	opts := []Options{
		WithMock(mockRabbitMQ),
	}

	service, err := NewService(opts...)
	assert.NoError(t, err)
	assert.NotNil(t, service)

	assert.NotNil(t, service.MonitorConn())
}

func TestConsume(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCh := NewMockChannelRepo(ctrl)
	mockConn := NewMockConnectionRepo(ctrl)
	mockRabbitMQ := NewMockRabbitMQRepo(ctrl)

	mockCh.EXPECT().Consume(gomock.Any(), gomock.Any(),
		gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, nil).AnyTimes()
	mockRabbitMQ.EXPECT().GetCh().Return(mockCh).AnyTimes()
	mockRabbitMQ.EXPECT().GetConn().Return(mockConn).AnyTimes()

	opts := []Options{
		WithMock(mockRabbitMQ),
	}

	service, err := NewService(opts...)
	assert.NoError(t, err)
	assert.NotNil(t, service)

	_, err = service.Consume("teste", "teste",
		true, true, true, true, nil)
	assert.NoError(t, err)
}

func TestPublish(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCh := NewMockChannelRepo(ctrl)
	mockConn := NewMockConnectionRepo(ctrl)
	mockRabbitMQ := NewMockRabbitMQRepo(ctrl)

	mockCh.EXPECT().Publish(gomock.Any(), gomock.Any(),
		gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
	mockRabbitMQ.EXPECT().GetCh().Return(mockCh).AnyTimes()
	mockRabbitMQ.EXPECT().GetConn().Return(mockConn).AnyTimes()

	opts := []Options{
		WithMock(mockRabbitMQ),
	}

	service, err := NewService(opts...)
	assert.NoError(t, err)
	assert.NotNil(t, service)

	err = service.Publish("teste", "teste", []byte("teste"))
	assert.NoError(t, err)

}

func TestQueueDeclare(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCh := NewMockChannelRepo(ctrl)
	mockConn := NewMockConnectionRepo(ctrl)
	mockRabbitMQ := NewMockRabbitMQRepo(ctrl)

	mockCh.EXPECT().QueueDeclare(gomock.Any(), gomock.Any(),
		gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(amqp.Queue{}, nil).AnyTimes()
	mockCh.EXPECT().QueueBind(gomock.Any(), gomock.Any(),
		gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
	mockRabbitMQ.EXPECT().GetCh().Return(mockCh).AnyTimes()
	mockRabbitMQ.EXPECT().GetConn().Return(mockConn).AnyTimes()

	opts := []Options{
		WithMock(mockRabbitMQ),
	}

	service, err := NewService(opts...)
	assert.NoError(t, err)
	assert.NotNil(t, service)

	_, err = service.QueueDeclare("teste", true,
		true, true, true, nil)
	assert.NoError(t, err)
}

func TestExchangeDeclare(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCh := NewMockChannelRepo(ctrl)
	mockConn := NewMockConnectionRepo(ctrl)
	mockRabbitMQ := NewMockRabbitMQRepo(ctrl)

	mockCh.EXPECT().ExchangeDeclare(gomock.Any(), gomock.Any(),
		gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
	mockRabbitMQ.EXPECT().GetCh().Return(mockCh).AnyTimes()
	mockRabbitMQ.EXPECT().GetConn().Return(mockConn).AnyTimes()

	opts := []Options{
		WithMock(mockRabbitMQ),
	}

	service, err := NewService(opts...)
	assert.NoError(t, err)
	assert.NotNil(t, service)

	err = service.ExchangeDeclare("teste", "teste",
		true, true, true, true, nil)
	assert.NoError(t, err)
}

func TestQueueBind(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCh := NewMockChannelRepo(ctrl)
	mockConn := NewMockConnectionRepo(ctrl)
	mockRabbitMQ := NewMockRabbitMQRepo(ctrl)

	mockCh.EXPECT().QueueBind(gomock.Any(), gomock.Any(),
		gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
	mockRabbitMQ.EXPECT().GetCh().Return(mockCh).AnyTimes()
	mockRabbitMQ.EXPECT().GetConn().Return(mockConn).AnyTimes()

	opts := []Options{
		WithMock(mockRabbitMQ),
	}

	service, err := NewService(opts...)
	assert.NoError(t, err)
	assert.NotNil(t, service)

	err = service.QueueBind("teste", "teste",
		"teste", true, nil)
	assert.NoError(t, err)
}

func TestQueueUnbind(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCh := NewMockChannelRepo(ctrl)
	mockConn := NewMockConnectionRepo(ctrl)
	mockRabbitMQ := NewMockRabbitMQRepo(ctrl)

	mockCh.EXPECT().QueueUnbind(gomock.Any(), gomock.Any(),
		gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
	mockRabbitMQ.EXPECT().GetCh().Return(mockCh).AnyTimes()
	mockRabbitMQ.EXPECT().GetConn().Return(mockConn).AnyTimes()

	opts := []Options{
		WithMock(mockRabbitMQ),
	}

	service, err := NewService(opts...)
	assert.NoError(t, err)
	assert.NotNil(t, service)

	err = service.QueueUnbind("teste", "teste",
		"teste", nil)
	assert.NoError(t, err)
}

func TestQueueDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCh := NewMockChannelRepo(ctrl)
	mockConn := NewMockConnectionRepo(ctrl)
	mockRabbitMQ := NewMockRabbitMQRepo(ctrl)

	mockCh.EXPECT().QueueDelete(gomock.Any(), gomock.Any(),
		gomock.Any(), gomock.Any()).Return(0, nil).AnyTimes()
	mockRabbitMQ.EXPECT().GetCh().Return(mockCh).AnyTimes()
	mockRabbitMQ.EXPECT().GetConn().Return(mockConn).AnyTimes()

	opts := []Options{
		WithMock(mockRabbitMQ),
	}

	service, err := NewService(opts...)
	assert.NoError(t, err)
	assert.NotNil(t, service)

	_, err = service.QueueDelete("teste", true, true, true)
	assert.NoError(t, err)
}

func TestExchangeDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCh := NewMockChannelRepo(ctrl)
	mockConn := NewMockConnectionRepo(ctrl)
	mockRabbitMQ := NewMockRabbitMQRepo(ctrl)

	mockCh.EXPECT().ExchangeDelete(gomock.Any(), gomock.Any(),
		gomock.Any()).Return(nil).AnyTimes()
	mockRabbitMQ.EXPECT().GetCh().Return(mockCh).AnyTimes()
	mockRabbitMQ.EXPECT().GetConn().Return(mockConn).AnyTimes()

	opts := []Options{
		WithMock(mockRabbitMQ),
	}

	service, err := NewService(opts...)
	assert.NoError(t, err)
	assert.NotNil(t, service)

	err = service.ExchangeDelete("teste", true, true)
	assert.NoError(t, err)
}
