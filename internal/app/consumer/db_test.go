package consumer

import (
	"github.com/golang/mock/gomock"
	"github.com/ozonmp/bss-office-api/internal/mocks"
	"github.com/ozonmp/bss-office-api/internal/model"
	"testing"
	"time"
)

func Test_consumer_Start(t *testing.T) {
	events := make(chan model.OfficeEvent, 512)

	ctrl := gomock.NewController(t)
	repo := mocks.NewMockEventRepo(ctrl)

	repo.EXPECT().Lock(gomock.Any()).
		DoAndReturn(func(batchSize uint64) ([]model.OfficeEvent, error) {
			result := make([]model.OfficeEvent, batchSize)
			for i := uint64(0); i < batchSize; i++ {
				result = append(result, model.OfficeEvent{
					ID:     i,
					Type:   model.Created,
					Status: model.Deferred,
					Entity: &model.Office{},
				})
			}
			return result, nil
		}).AnyTimes()

	cfg := Config{
		n:         2,
		events:    events,
		repo:      repo,
		batchSize: 10,
		timeout:   time.Second * 1,
	}

	consumer := NewDbConsumer(
		cfg.n,
		cfg.batchSize,
		cfg.timeout,
		cfg.repo,
		cfg.events)

	consumer.Start()

	<-events
	<-events
	<-events
}
