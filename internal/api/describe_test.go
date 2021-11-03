package api

import (
	"context"
	bss_office_api "github.com/ozonmp/bss-office-api/pkg/bss-office-api"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
	"time"
)

func Test_officeAPI_DescribeOfficeV1(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		request  *bss_office_api.DescribeOfficeV1Request
		wantCode codes.Code
		wantMsg  string
	}{
		{
			name: "with zero id",
			request: &bss_office_api.DescribeOfficeV1Request{
				OfficeId: 0,
			},
			wantCode: codes.InvalidArgument,
			wantMsg:  "invalid DescribeOfficeV1Request.OfficeId: value must be greater than 0",
		},
		{
			name:     "with without id",
			request:  &bss_office_api.DescribeOfficeV1Request{},
			wantCode: codes.InvalidArgument,
			wantMsg:  "invalid DescribeOfficeV1Request.OfficeId: value must be greater than 0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fixture := setUp(t)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			_, err := fixture.apiServer.DescribeOfficeV1(ctx, tt.request)

			actualStatus, _ := status.FromError(err)

			assert.Equal(t, tt.wantCode, actualStatus.Code())
			assert.Equal(t, tt.wantMsg, actualStatus.Message())
		})
	}
}
