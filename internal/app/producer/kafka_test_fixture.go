package producer

import (
	"github.com/golang/mock/gomock"
	"github.com/ozonmp/bss-office-api/internal/mocks"
	"testing"
)

type ProducerFixture struct {
	Repo   *mocks.MockEventRepo
	Sender *mocks.MockEventSender
	ctrl   *gomock.Controller
}

func LoadFixture(t *testing.T) ProducerFixture {
	ctrl := gomock.NewController(t)
	repo := mocks.NewMockEventRepo(ctrl)
	sender := mocks.NewMockEventSender(ctrl)

	return ProducerFixture{
		Repo:   repo,
		Sender: sender,
		ctrl:   ctrl,
	}
}
