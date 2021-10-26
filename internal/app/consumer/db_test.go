package consumer

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/ozonmp/bss-office-api/internal/mocks"
	"github.com/ozonmp/bss-office-api/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

const testBatchSize = uint64(10)
const testConsumerCount = 2
const testEventBufferSize = 512

type ConsumerTestSuite struct {
	suite.Suite
	consumer Consumer
	repo     *mocks.MockEventRepo
	ctrl     *gomock.Controller
	model    model.OfficeEvent
	events   chan model.OfficeEvent
}

func (suite *ConsumerTestSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.repo = mocks.NewMockEventRepo(suite.ctrl)
	suite.events = make(chan model.OfficeEvent, testEventBufferSize)

	suite.consumer = NewDbConsumer(
		testConsumerCount,
		testBatchSize,
		time.Millisecond,
		suite.repo,
		suite.events,
	)

	suite.model = model.OfficeEvent{
		ID:     1,
		Type:   model.Created,
		Status: model.Deferred,
		Entity: &model.Office{},
	}
}

func TestConsumerTestSuite(t *testing.T) {
	suite.Run(t, new(ConsumerTestSuite))
}

func (suite *ConsumerTestSuite) Test_consumer_Start() {
	suite.repo.EXPECT().Lock(gomock.Eq(testBatchSize)).Return([]model.OfficeEvent{suite.model}, nil).MinTimes(1)

	suite.consumer.Start()
	defer suite.consumer.Close()

	time.Sleep(time.Millisecond * 2)

	timer := time.NewTimer(time.Second)

	select {
	case event, ok := <-suite.events:
		if !ok {
			suite.Fail("cannot get event from the channel")
		}
		assert.Equal(suite.T(), event, suite.model)
		timer.Stop()

	case <-timer.C:
		suite.Fail("timeout waiting event")
	}
}

func (suite *ConsumerTestSuite) Test_consumer_Error() {
	suite.repo.EXPECT().Lock(gomock.Eq(testBatchSize)).
		Return([]model.OfficeEvent{suite.model, suite.model}, errors.New("test lock error")).MinTimes(1)

	suite.consumer.Start()
	defer suite.consumer.Close()

	time.Sleep(time.Millisecond * 2)
	assert.Len(suite.T(), suite.events, 0)
}
