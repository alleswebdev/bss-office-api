package producer

import (
	"context"
	"errors"
	"github.com/gammazero/workerpool"
	"github.com/golang/mock/gomock"
	"github.com/ozonmp/bss-office-api/internal/mocks"
	"github.com/ozonmp/bss-office-api/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestProducer_Start(t *testing.T) {
	t.Parallel()
	events := make(chan model.OfficeEvent, 10)

	ctrl := gomock.NewController(t)
	repo := mocks.NewMockEventRepo(ctrl)
	sender := mocks.NewMockEventSender(ctrl)
	workerPool := workerpool.New(5)

	producer := NewKafkaProducer(
		2,
		sender,
		repo,
		events,
		workerPool)

	ctx, cancel := context.WithCancel(context.Background())

	producer.Start(ctx)
	cancel()
	producer.Close()
}

func TestProducer_With_Update(t *testing.T) {
	t.Parallel()
	events := make(chan model.OfficeEvent, 1)

	ctrl := gomock.NewController(t)
	repo := mocks.NewMockEventRepo(ctrl)
	sender := mocks.NewMockEventSender(ctrl)

	repo.EXPECT().Unlock(gomock.Any()).Return(nil).MaxTimes(0)
	repo.EXPECT().Remove(gomock.Any()).Return(nil).MinTimes(1)
	sender.EXPECT().Send(gomock.Any(), gomock.Any()).Return(nil).MinTimes(1)

	workerPool := workerpool.New(5)
	defer workerPool.StopWait()

	producer := NewKafkaProducer(
		1,
		sender,
		repo,
		events,
		workerPool)

	events <- model.OfficeEvent{
		ID:     1,
		Type:   model.Created,
		Status: model.Deferred,
		Entity: nil,
	}

	ctx, cancel := context.WithCancel(context.Background())
	producer.Start(ctx)

	time.Sleep(time.Millisecond)
	assert.Len(t, events, 0)

	cancel()
	producer.Close()
}

func TestProducer_With_Error(t *testing.T) {
	t.Parallel()
	events := make(chan model.OfficeEvent, 1)

	ctrl := gomock.NewController(t)
	repo := mocks.NewMockEventRepo(ctrl)
	sender := mocks.NewMockEventSender(ctrl)

	repo.EXPECT().Unlock(gomock.Any()).Return(nil).MinTimes(1)
	repo.EXPECT().Remove(gomock.Any()).Return(nil).MaxTimes(0)
	sender.EXPECT().Send(gomock.Any(), gomock.Any()).Return(errors.New("error sending")).MinTimes(1)

	workerPool := workerpool.New(5)
	defer workerPool.StopWait()

	producer := NewKafkaProducer(
		1,
		sender,
		repo,
		events,
		workerPool)

	events <- model.OfficeEvent{
		ID:     1,
		Type:   model.Created,
		Status: model.Deferred,
		Entity: nil,
	}

	ctx, cancel := context.WithCancel(context.Background())
	producer.Start(ctx)
	defer producer.Close()
	defer cancel()

	time.Sleep(time.Millisecond)
	assert.Len(t, events, 0)
}
