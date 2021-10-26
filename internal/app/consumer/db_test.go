package consumer

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/ozonmp/bss-office-api/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const testBatchSize = uint64(10)
const testConsumerCount = 2
const testEventBufferSize = 512

var testModel = model.OfficeEvent{
	ID:     1,
	Type:   model.Created,
	Status: model.Deferred,
	Entity: &model.Office{},
}

func Test_consumer_Start(t *testing.T) {
	t.Parallel()

	fixture := LoadFixture(t)
	fixture.Repo.EXPECT().Lock(gomock.Eq(testBatchSize)).Return([]model.OfficeEvent{testModel}, nil).MinTimes(1)

	events := make(chan model.OfficeEvent, testEventBufferSize)
	consumer := NewDbConsumer(
		testConsumerCount,
		testBatchSize,
		time.Millisecond,
		fixture.Repo,
		events,
	)

	consumer.Start()
	defer consumer.Close()

	time.Sleep(time.Millisecond * 2)

	timer := time.NewTimer(time.Second)

	select {
	case event, ok := <-events:
		if !ok {
			t.Fatal("cannot get event from the channel")
		}
		assert.Equal(t, event, testModel)
		timer.Stop()

	case <-timer.C:
		t.Fatal("timeout waiting event")
	}
}

func Test_consumer_Error(t *testing.T) {
	t.Parallel()
	events := make(chan model.OfficeEvent, testEventBufferSize)

	fixture := LoadFixture(t)

	fixture.Repo.EXPECT().Lock(gomock.Eq(testBatchSize)).
		Return([]model.OfficeEvent{testModel, testModel}, errors.New("test lock error")).MinTimes(1)

	consumer := NewDbConsumer(
		testConsumerCount,
		testBatchSize,
		time.Millisecond,
		fixture.Repo,
		events,
	)

	consumer.Start()
	defer consumer.Close()

	time.Sleep(time.Millisecond * 2)
	assert.Len(t, events, 0)
}
