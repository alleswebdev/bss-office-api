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

<<<<<<< HEAD
	repo.EXPECT().Unlock(gomock.Any()).Return(nil).MaxTimes(0)
	repo.EXPECT().Remove(gomock.Any()).Return(nil).MinTimes(1)
	sender.EXPECT().Send(gomock.Any(), gomock.Any()).Return(nil).MinTimes(1)
=======
	fixture.Repo.EXPECT().Unlock(gomock.Any()).Return(nil).Times(0)
	fixture.Repo.EXPECT().Remove(gomock.Eq([]uint64{testModel.ID})).Return(nil).Times(1).After(
		fixture.Sender.EXPECT().Send(gomock.Eq(&testModel)).Return(nil).Times(1))
>>>>>>> test: expanded test cases

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

	ctx, cancel := context.WithCancel(context.Background())
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

	fixture.Repo.EXPECT().Remove(gomock.Any()).Return(nil).Times(0)

<<<<<<< HEAD
	repo.EXPECT().Unlock(gomock.Any()).Return(nil).MinTimes(1)
	repo.EXPECT().Remove(gomock.Any()).Return(nil).MaxTimes(0)
	sender.EXPECT().Send(gomock.Any(), gomock.Any()).Return(errors.New("error sending")).MinTimes(1)
=======
	fixture.Repo.EXPECT().Unlock(gomock.Eq([]uint64{testModel.ID})).Return(nil).Times(1).After(
		fixture.Sender.EXPECT().Send(gomock.Eq(&testModel)).Return(errors.New("test error")).Times(1))
>>>>>>> test: expanded test cases

	workerPool := workerpool.New(5)
	defer workerPool.StopWait()

	producer := NewKafkaProducer(
		testProducerCount,
		fixture.Sender,
		fixture.Repo,
		events,
		workerPool)

	events <- testModel

	ctx, cancel := context.WithCancel(context.Background())
	producer.Start(ctx)
	defer producer.Close()
	defer cancel()

	time.Sleep(time.Millisecond)
	assert.Len(t, events, 0)
}
