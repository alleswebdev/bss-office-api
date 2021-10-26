package producer

import (
	"context"
	"errors"
	"github.com/gammazero/workerpool"
	"github.com/golang/mock/gomock"
	"github.com/ozonmp/bss-office-api/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const testProducerCount = 2
const testWorkerCount = 2
const testEventBufferSize = 512

var testModel = model.OfficeEvent{
	ID:     1,
	Type:   model.Created,
	Status: model.Deferred,
	Entity: &model.Office{},
}

func TestProducer_Start(t *testing.T) {
	t.Parallel()

	events := make(chan model.OfficeEvent, testEventBufferSize)
	defer close(events)

	fixture := LoadFixture(t)
	workerPool := workerpool.New(testWorkerCount)

	producer := NewKafkaProducer(
		testProducerCount,
		fixture.Sender,
		fixture.Repo,
		events,
		workerPool)

	ctx, cancel := context.WithCancel(context.Background())

	producer.Start(ctx)
	cancel()
	producer.Close()
}

func TestProducer_Update(t *testing.T) {
	t.Parallel()
	events := make(chan model.OfficeEvent, testEventBufferSize)
	defer close(events)

	fixture := LoadFixture(t)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	fixture.Repo.EXPECT().Unlock(gomock.Any()).Return(nil).Times(0)

	fixture.Repo.EXPECT().Remove(gomock.Eq([]uint64{testModel.ID})).Return(nil).Times(1).After(
		fixture.Sender.EXPECT().Send(gomock.Eq(ctx), gomock.Eq(&testModel)).Return(nil).Times(1))

	workerPool := workerpool.New(testWorkerCount)
	defer workerPool.StopWait()

	producer := NewKafkaProducer(
		testProducerCount,
		fixture.Sender,
		fixture.Repo,
		events,
		workerPool)

	events <- testModel

	assert.Len(t, events, 1)

	producer.Start(ctx)

	time.Sleep(time.Millisecond * 5)
	assert.Len(t, events, 0)

	cancel()
	producer.Close()
}

func TestProducer_With_Error(t *testing.T) {
	t.Parallel()
	events := make(chan model.OfficeEvent, testEventBufferSize)
	defer close(events)

	fixture := LoadFixture(t)

	ctx, cancel := context.WithCancel(context.Background())

	fixture.Repo.EXPECT().Remove(gomock.Any()).Return(nil).Times(0)

	fixture.Repo.EXPECT().Unlock(gomock.Eq([]uint64{testModel.ID})).Return(nil).Times(1).After(
		fixture.Sender.EXPECT().Send(gomock.Eq(ctx), gomock.Eq(&testModel)).Return(errors.New("test error")).Times(1))

	workerPool := workerpool.New(5)
	defer workerPool.StopWait()

	producer := NewKafkaProducer(
		testProducerCount,
		fixture.Sender,
		fixture.Repo,
		events,
		workerPool)

	events <- testModel

	producer.Start(ctx)

	time.Sleep(time.Millisecond)
	assert.Len(t, events, 0)

	cancel()
	producer.Close()
}
