package consumer

import (
	"github.com/golang/mock/gomock"
	"github.com/ozonmp/bss-office-api/internal/mocks"
	"testing"
)

type ConsumerFixture struct {
	Repo *mocks.MockEventRepo
	ctrl *gomock.Controller
}

func LoadFixture(t *testing.T) ConsumerFixture {
	ctrl := gomock.NewController(t)
	repo := mocks.NewMockEventRepo(ctrl)

	return ConsumerFixture{
		Repo: repo,
		ctrl: ctrl,
	}
}
