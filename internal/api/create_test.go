package api

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/ozonmp/bss-office-api/internal/mocks"
	bss_office_api "github.com/ozonmp/bss-office-api/pkg/bss-office-api"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
	"time"
)

type ApiFixture struct {
	repo      *mocks.MockRepo
	ctrl      *gomock.Controller
	apiServer bss_office_api.BssOfficeApiServiceServer
}

func setUp(t *testing.T) ApiFixture {
	var fixture ApiFixture

	fixture.ctrl = gomock.NewController(t)
	fixture.repo = mocks.NewMockRepo(fixture.ctrl)
	fixture.apiServer = NewOfficeAPI(fixture.repo)

	return fixture
}

func Test_officeAPI_CreateOfficeV1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		request  *bss_office_api.CreateOfficeV1Request
		wantCode codes.Code
		wantMsg  string
	}{
		{
			name:     "with empty name",
			request:  &bss_office_api.CreateOfficeV1Request{Name: ""},
			wantCode: codes.InvalidArgument,
			wantMsg:  "invalid CreateOfficeV1Request.Name: value length must be between 2 and 100 runes, inclusive",
		},
		{
			name:     "with one rune name",
			request:  &bss_office_api.CreateOfficeV1Request{Name: "a"},
			wantCode: codes.InvalidArgument,
			wantMsg:  "invalid CreateOfficeV1Request.Name: value length must be between 2 and 100 runes, inclusive",
		},
		{
			name:     "with 101 rune name",
			request:  &bss_office_api.CreateOfficeV1Request{Name: "Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Aenean commodo ligula eget dolor. Aenean mass"},
			wantCode: codes.InvalidArgument,
			wantMsg:  "invalid CreateOfficeV1Request.Name: value length must be between 2 and 100 runes, inclusive",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fixture := setUp(t)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			_, err := fixture.apiServer.CreateOfficeV1(ctx, tt.request)

			actualStatus, _ := status.FromError(err)

			assert.Equal(t, tt.wantCode, actualStatus.Code())
			assert.Equal(t, tt.wantMsg, actualStatus.Message())
		})
	}
}
