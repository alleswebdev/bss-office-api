package producer

import (
	"context"
	"errors"
	"fmt"
	"github.com/gammazero/workerpool"
	"github.com/golang/mock/gomock"
	"github.com/ozonmp/bss-office-api/internal/mocks"
	"github.com/ozonmp/bss-office-api/internal/model"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

const testProducerCount = 1
const testWorkerCount = 2
const testEventBufferSize = 512
const testBatchSize = 2
const testTimeout = time.Millisecond * 5

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
		testBatchSize,
		testTimeout,
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	gomock.InOrder(
		fixture.sender.EXPECT().Send(gomock.Any(), gomock.Eq(&fixture.model)).Return(nil),
		fixture.repo.EXPECT().Remove(gomock.Eq([]uint64{fixture.model.ID})).DoAndReturn(func(eventIDs []uint64) error {
			wg.Done()
			return nil
		}),
	)

	fixture.events <- fixture.model
	assert.Len(t, fixture.events, 1)

	fixture.producer.Start(ctx)
	defer fixture.producer.Close()

	wg.Wait()
	assert.Len(t, fixture.events, 0)

	cancel()
}

func TestProducer_With_Error(t *testing.T) {
	t.Parallel()

	fixture := setUp(t)
	defer fixture.tearDown()

	var wg sync.WaitGroup
	wg.Add(1)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	gomock.InOrder(
		fixture.sender.EXPECT().Send(gomock.Any(), gomock.Eq(&fixture.model)).Return(errors.New("test error")).Times(1),
		fixture.repo.EXPECT().Unlock(gomock.Eq([]uint64{fixture.model.ID})).DoAndReturn(func(eventIDs []uint64) error {
			wg.Done()
			return nil
		}).Times(1),
	)

	fixture.events <- fixture.model

	fixture.producer.Start(ctx)
	defer fixture.producer.Close()

	wg.Wait()
	assert.Len(t, fixture.events, 0)

	cancel()
}

// проверяем, что при пакетном анлоке буфер отсылается при заполнении до batchSize
func TestProducer_Batch_Error(t *testing.T) {
	t.Parallel()

	fixture := setUp(t)
	defer fixture.tearDown()

	var wg sync.WaitGroup
	wg.Add(2)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	fixture.sender.EXPECT().Send(gomock.Any(), gomock.Eq(&fixture.model)).Return(errors.New("test error")).Times(4)
	// batch size = 2
	fixture.repo.EXPECT().Unlock(gomock.Any()).DoAndReturn(func(eventIDs []uint64) error {
		defer wg.Done()
		assert.Equal(t, []uint64{fixture.model.ID, fixture.model.ID}, eventIDs)
		return nil
	}).Times(2)

	fixture.events <- fixture.model
	fixture.events <- fixture.model
	fixture.events <- fixture.model
	fixture.events <- fixture.model

	fixture.producer.StartBatch(ctx)
	defer fixture.producer.Close()

	wg.Wait()
	assert.Len(t, fixture.events, 0)
	cancel()
}

// Проверяем, что при пакетном анлоке незаполненный буфер отсылается по таймауту
func TestProducer_Batch_Error_Timeout(t *testing.T) {
	t.Parallel()

	fixture := setUp(t)
	defer fixture.tearDown()

	var wg sync.WaitGroup
	wg.Add(1)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	fixture.sender.EXPECT().Send(gomock.Any(), gomock.Eq(&fixture.model)).Return(errors.New("test error"))
	fixture.repo.EXPECT().Unlock(gomock.Any()).DoAndReturn(func(eventIDs []uint64) error {
		defer wg.Done()
		assert.Equal(t, []uint64{fixture.model.ID}, eventIDs)
		return nil
	})

	fixture.events <- fixture.model

	fixture.producer.StartBatch(ctx)
	defer fixture.producer.Close()

	wg.Wait()
	assert.Len(t, fixture.events, 0)
	cancel()
}

// проверяем, что при пакетном удалении буфер отсылается при заполнении до batchSize
func TestProducer_Batch_Start(t *testing.T) {
	t.Parallel()

	fixture := setUp(t)
	defer fixture.tearDown()

	var wg sync.WaitGroup
	wg.Add(2)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	fixture.sender.EXPECT().Send(gomock.Any(), gomock.Eq(&fixture.model)).Return(nil).Times(4)
	// batch size = 2
	fixture.repo.EXPECT().Remove(gomock.Any()).DoAndReturn(func(eventIDs []uint64) error {
		defer wg.Done()
		assert.Equal(t, []uint64{fixture.model.ID, fixture.model.ID}, eventIDs)
		return nil
	}).Times(2)

	fixture.events <- fixture.model
	fixture.events <- fixture.model
	fixture.events <- fixture.model
	fixture.events <- fixture.model

	fixture.producer.StartBatch(ctx)
	defer fixture.producer.Close()

	wg.Wait()
	assert.Len(t, fixture.events, 0)
	cancel()
}

// Проверяем, что при пакетном удалении незаполненный буфер отсылается по таймауту
func TestProducer_Batch_Start_Timeout(t *testing.T) {
	t.Parallel()

	fixture := setUp(t)
	defer fixture.tearDown()

	var wg sync.WaitGroup
	wg.Add(1)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	fixture.sender.EXPECT().Send(gomock.Any(), gomock.Eq(&fixture.model)).Return(nil)
	fixture.repo.EXPECT().Remove(gomock.Any()).DoAndReturn(func(eventIDs []uint64) error {
		defer wg.Done()
		assert.Equal(t, []uint64{fixture.model.ID}, eventIDs)
		return nil
	})

	fixture.events <- fixture.model

	fixture.producer.StartBatch(ctx)
	defer fixture.producer.Close()

	wg.Wait()
	assert.Len(t, fixture.events, 0)
	cancel()
}

// проверим, что при ошибке в методе удаления репозитория
// событие не потеряется и переотправится
func TestProducer_Batch_Start_With_Remove_err(t *testing.T) {
	t.Parallel()

	fixture := setUp(t)
	defer fixture.tearDown()

	var wg sync.WaitGroup
	wg.Add(2)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	fixture.sender.EXPECT().Send(gomock.Any(), gomock.Eq(&fixture.model)).Return(nil)

	fixture.repo.EXPECT().Remove(gomock.Any()).DoAndReturn(func(eventIDs []uint64) error {
		defer wg.Done()
		assert.Equal(t, []uint64{fixture.model.ID}, eventIDs)
		return fmt.Errorf("test error in repo")
	})

	fixture.repo.EXPECT().Remove(gomock.Any()).DoAndReturn(func(eventIDs []uint64) error {
		defer wg.Done()
		assert.Equal(t, []uint64{fixture.model.ID}, eventIDs)
		return nil
	})

	fixture.events <- fixture.model

	fixture.producer.StartBatch(ctx)
	defer fixture.producer.Close()

	wg.Wait()
	assert.Len(t, fixture.events, 0)
	cancel()
	fmt.Println("done")
}
