package producer

import (
	"context"
	"errors"
	"github.com/gammazero/workerpool"
	"github.com/golang/mock/gomock"
	"github.com/ozonmp/bss-office-api/internal/mocks"
	"github.com/ozonmp/bss-office-api/internal/model"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

const testProducerCount = 2
const testWorkerCount = 2
const testEventBufferSize = 512

type ProducerFixture struct {
	producer   Producer
	repo       *mocks.MockEventRepo
	ctrl       *gomock.Controller
	sender     *mocks.MockEventSender
	model      model.OfficeEvent
	events     chan model.OfficeEvent
	workerPool *workerpool.WorkerPool
}

func setUp(t *testing.T) ProducerFixture {
	var fixture ProducerFixture

	fixture.ctrl = gomock.NewController(t)
	fixture.repo = mocks.NewMockEventRepo(fixture.ctrl)
	fixture.sender = mocks.NewMockEventSender(fixture.ctrl)
	fixture.events = make(chan model.OfficeEvent, testEventBufferSize)
	fixture.workerPool = workerpool.New(testWorkerCount)
	fixture.producer = NewKafkaProducer(
		testProducerCount,
		fixture.sender,
		fixture.repo,
		fixture.events,
		fixture.workerPool)

	fixture.model = model.OfficeEvent{
		ID:     1,
		Type:   model.Created,
		Status: model.Deferred,
		Entity: &model.Office{},
	}

	return fixture
}

func (f *ProducerFixture) tearDown() {
	f.ctrl.Finish()
	close(f.events)
	f.workerPool.Stop()
}

func TestProducer_Update(t *testing.T) {
	t.Parallel()

	fixture := setUp(t)
	defer fixture.tearDown()

	var wg sync.WaitGroup
	wg.Add(1)

	gomock.InOrder(
		fixture.sender.EXPECT().Send(gomock.Eq(&fixture.model)).Return(nil).Times(1),
		fixture.repo.EXPECT().Remove(gomock.Eq([]uint64{fixture.model.ID})).DoAndReturn(func(eventIDs []uint64) error {
			wg.Done()
			return nil
		}).Times(1),
	)

	fixture.events <- fixture.model
	assert.Len(t, fixture.events, 1)

	fixture.producer.Start()
	defer fixture.producer.Close()

	wg.Wait()
	assert.Len(t, fixture.events, 0)
}

func TestProducer_With_Error(t *testing.T) {
	t.Parallel()

	fixture := setUp(t)
	defer fixture.tearDown()

	var wg sync.WaitGroup
	wg.Add(1)

	gomock.InOrder(
		fixture.sender.EXPECT().Send(gomock.Eq(&fixture.model)).Return(errors.New("test error")).Times(1),
		fixture.repo.EXPECT().Unlock(gomock.Eq([]uint64{fixture.model.ID})).DoAndReturn(func(eventIDs []uint64) error {
			wg.Done()
			return nil
		}).Times(1),
	)

	fixture.events <- fixture.model

	fixture.producer.Start()
	defer fixture.producer.Close()

	wg.Wait()
	assert.Len(t, fixture.events, 0)
}
