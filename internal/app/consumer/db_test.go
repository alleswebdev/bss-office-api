package consumer

import (
	"github.com/golang/mock/gomock"
	"github.com/ozonmp/bss-office-api/internal/mocks"
	"github.com/ozonmp/bss-office-api/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_consumer_Start(t *testing.T) {
	t.Parallel()
	events := make(chan model.OfficeEvent, 10)

	ctrl := gomock.NewController(t)
	repo := mocks.NewMockEventRepo(ctrl)

	testModel := model.OfficeEvent{
		ID:     1,
		Type:   model.Created,
		Status: model.Deferred,
		Entity: &model.Office{},
	}

	repo.EXPECT().Lock(gomock.Any()).Return([]model.OfficeEvent{testModel}, nil).MinTimes(1)

	cfg := Config{
		n:         2,
		events:    events,
		repo:      repo,
		batchSize: 10,
		timeout:   time.Millisecond * 1,
	}

	consumer := NewDbConsumer(
		cfg.n,
		cfg.batchSize,
		cfg.timeout,
		cfg.repo,
		cfg.events)

	consumer.Start()
	defer consumer.Close()

	time.Sleep(time.Millisecond * 2)

	timer := time.NewTimer(time.Second)
	for {
		select {
		case event, ok := <-events:
			if !ok {
				t.Fatal("cannot get event from the channel")
			}
			assert.Equal(t, event, testModel)
			timer.Stop()
			return

		case <-timer.C:
			t.Fatal("timeout waiting event")
		}
	}

}
