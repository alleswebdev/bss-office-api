package producer

import (
	"context"
	"errors"
	"github.com/gammazero/workerpool"
	"github.com/golang/mock/gomock"
	"github.com/ozonmp/bss-office-api/internal/mocks"
	"github.com/ozonmp/bss-office-api/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

const testProducerCount = 2
const testWorkerCount = 2
const testEventBufferSize = 512

type ProducerTestSuite struct {
	suite.Suite
	producer   Producer
	repo       *mocks.MockEventRepo
	ctrl       *gomock.Controller
	sender     *mocks.MockEventSender
	model      model.OfficeEvent
	events     chan model.OfficeEvent
	workerPool *workerpool.WorkerPool
}

func (s *ProducerTestSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.repo = mocks.NewMockEventRepo(s.ctrl)
	s.sender = mocks.NewMockEventSender(s.ctrl)
	s.events = make(chan model.OfficeEvent, testEventBufferSize)
	s.workerPool = workerpool.New(testWorkerCount)
	s.producer = NewKafkaProducer(
		testProducerCount,
		s.sender,
		s.repo,
		s.events,
		s.workerPool)

	s.model = model.OfficeEvent{
		ID:     1,
		Type:   model.Created,
		Status: model.Deferred,
		Entity: &model.Office{},
	}
}

func (s *ProducerTestSuite) TearDownTest() {
	close(s.events)
	s.workerPool.Stop()
}

func TestProducerTestSuite(t *testing.T) {
	suite.Run(t, new(ProducerTestSuite))
}

func (s *ProducerTestSuite) TestProducer_Start() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	s.producer.Start(ctx)
	s.producer.Close()
	cancel()
}

func (s *ProducerTestSuite) TestProducer_Update() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	s.repo.EXPECT().Unlock(gomock.Any()).Return(nil).Times(0)
	s.repo.EXPECT().Remove(gomock.Eq([]uint64{s.model.ID})).Return(nil).Times(1).After(
		s.sender.EXPECT().Send(gomock.Eq(ctx), gomock.Eq(&s.model)).Return(nil).Times(1))

	s.events <- s.model

	assert.Len(s.T(), s.events, 1)

	s.producer.Start(ctx)
	defer s.producer.Close()

	time.Sleep(time.Millisecond * 5)
	assert.Len(s.T(), s.events, 0)

	cancel()
}

func (s *ProducerTestSuite) TestProducer_With_Error() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	s.repo.EXPECT().Remove(gomock.Any()).Return(nil).Times(0)

	s.repo.EXPECT().Unlock(gomock.Eq([]uint64{s.model.ID})).Return(nil).Times(1).After(
		s.sender.EXPECT().Send(gomock.Eq(ctx), gomock.Eq(&s.model)).Return(errors.New("test error")).Times(1))

	s.events <- s.model

	s.producer.Start(ctx)
	defer s.producer.Close()

	time.Sleep(time.Millisecond)
	assert.Len(s.T(), s.events, 0)

	cancel()
}
