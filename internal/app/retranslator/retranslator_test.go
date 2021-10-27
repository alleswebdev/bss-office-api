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

	ctx, cancel := context.WithCancel(context.Background())
	retranslator := NewRetranslator(cfg)
	retranslator.Start(ctx)
	cancel()
	repo.EXPECT().Lock(gomock.Eq(cfg.ConsumeSize)).Times(2)

	time.Sleep(time.Millisecond * 1) // убрать
	retranslator.Close()
}
