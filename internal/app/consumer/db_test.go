package consumer

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/ozonmp/bss-office-api/internal/mocks"
	"github.com/ozonmp/bss-office-api/internal/model"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

const testBatchSize = uint64(10)
const testConsumerCount = 1
const testEventBufferSize = 512

type ConsumerFixture struct {
	consumer Consumer
	repo     *mocks.MockEventRepo
	ctrl     *gomock.Controller
	model    model.OfficeEvent
	events   chan model.OfficeEvent
}

func setUp(t *testing.T) ConsumerFixture {
	var fixture ConsumerFixture
	fixture.ctrl = gomock.NewController(t)
	fixture.repo = mocks.NewMockEventRepo(fixture.ctrl)
	fixture.events = make(chan model.OfficeEvent, testEventBufferSize)

	fixture.consumer = NewDbConsumer(
		testConsumerCount,
		testBatchSize,
		time.Millisecond*10,
		fixture.repo,
		fixture.events,
	)

	fixture.model = model.OfficeEvent{
		ID:      1,
		Type:    model.Created,
		Status:  model.Deferred,
		Payload: model.OfficePayload{},
	}

	return fixture
}

func (f *ConsumerFixture) tearDown() {
	f.ctrl.Finish()
}

func Test_consumer_Start(t *testing.T) {
	t.Parallel()

	fixture := setUp(t)
	defer fixture.tearDown()

	fixture.repo.EXPECT().Lock(gomock.Any(), gomock.Eq(testBatchSize)).Return([]model.OfficeEvent{fixture.model}, nil).Times(testConsumerCount)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	fixture.consumer.Start(ctx)
	defer fixture.consumer.Close()

	timer := time.NewTimer(time.Second)

	select {
	case event, ok := <-fixture.events:
		if !ok {
			t.Error("cannot get event from the channel")
		}
		assert.Equal(t, event, fixture.model)
		timer.Stop()

	case <-timer.C:
		t.Error("timeout waiting event")
	}

	cancel()
}

func Test_consumer_Error(t *testing.T) {
	t.Parallel()

	fixture := setUp(t)
	defer fixture.tearDown()

	var wg sync.WaitGroup
	wg.Add(testConsumerCount)

	fixture.repo.EXPECT().Lock(gomock.Any(), gomock.Eq(testBatchSize)).DoAndReturn(func(ctx context.Context, n uint64) ([]model.OfficeEvent, error) {
		defer wg.Done()
		return []model.OfficeEvent{fixture.model, fixture.model}, errors.New("test lock error")
	}).Times(testConsumerCount)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	fixture.consumer.Start(ctx)
	defer fixture.consumer.Close()

	wg.Wait()
	assert.Len(t, fixture.events, 0)

	cancel()
}
