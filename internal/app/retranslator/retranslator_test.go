package retranslator

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/ozonmp/bss-office-api/internal/mocks"
)

func TestStart(t *testing.T) {

	ctrl := gomock.NewController(t)
	repo := mocks.NewMockEventRepo(ctrl)
	sender := mocks.NewMockEventSender(ctrl)

	cfg := Config{
		ChannelSize:    512,
		ConsumerCount:  2,
		ConsumeSize:    10,
		ConsumeTimeout: 1 * time.Millisecond,
		ProducerCount:  2,
		WorkerCount:    2,
		Repo:           repo,
		Sender:         sender,
	}

	repo.EXPECT().Lock(gomock.Eq(cfg.ConsumeSize)).Times(2)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	retranslator := NewRetranslator(cfg)
	retranslator.Start(ctx)

	time.Sleep(time.Millisecond * 1) // убрать

	cancel()
	retranslator.Close()
}
