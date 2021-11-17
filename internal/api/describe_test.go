package api

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/ozonmp/bss-office-api/internal/model"
	bss_office_api "github.com/ozonmp/bss-office-api/pkg/bss-office-api"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

const errDescribeOfficeIDValidation = "invalid DescribeOfficeV1Request.OfficeId: value must be greater than 0"

func Test_officeAPI_DescribeOfficeV1(t *testing.T) {
	t.Parallel()

	fixture := setUp(t)

	testOfficeID := uint64(1)

	fixture.officeRepo.EXPECT().DescribeOffice(gomock.Any(), testOfficeID).DoAndReturn(func(ctx context.Context, officeID uint64) (*model.Office, error) {
		return &model.Office{
			ID:          testOfficeID,
			Name:        "test",
			Description: "test",
		}, nil
	})

	res, err := fixture.apiServer.DescribeOfficeV1(context.Background(),
		&bss_office_api.DescribeOfficeV1Request{OfficeId: testOfficeID})

	assert.Equal(t, res.GetValue().GetId(), testOfficeID)
	assert.Equal(t, res.GetValue().GetName(), "test")
	assert.Equal(t, res.GetValue().GetDescription(), "test")
	assert.NoError(t, err)
}

func Test_officeAPI_DescribeOfficeV1_Not_Found(t *testing.T) {
	t.Parallel()

	fixture := setUp(t)

	testOfficeID := uint64(1)

	fixture.officeRepo.EXPECT().DescribeOffice(gomock.Any(), testOfficeID)

	_, err := fixture.apiServer.DescribeOfficeV1(context.Background(),
		&bss_office_api.DescribeOfficeV1Request{OfficeId: testOfficeID})

	actualStatus, _ := status.FromError(err)

	assert.Equal(t, codes.NotFound, actualStatus.Code())
	assert.Equal(t, "office not found", actualStatus.Message())
}

func Test_officeAPI_DescribeOfficeV1_Repo_Err(t *testing.T) {
	t.Parallel()

	fixture := setUp(t)

	testOfficeID := uint64(1)

	errTest := errors.New("test officeRepo err")
	fixture.officeRepo.EXPECT().DescribeOffice(gomock.Any(), testOfficeID).
		DoAndReturn(func(ctx context.Context, officeID uint64) (*model.Office, error) {
			return nil, errTest
		})

	_, err := fixture.apiServer.DescribeOfficeV1(context.Background(),
		&bss_office_api.DescribeOfficeV1Request{OfficeId: testOfficeID})

	actualStatus, _ := status.FromError(err)

	assert.Equal(t, codes.Internal, actualStatus.Code())
	assert.Error(t, errTest, actualStatus.Err())
}

func Test_officeAPI_DescribeOfficeV1_With_Zero_ID(t *testing.T) {
	t.Parallel()

	fixture := setUp(t)
	_, err := fixture.apiServer.DescribeOfficeV1(context.Background(),
		&bss_office_api.DescribeOfficeV1Request{OfficeId: 0})

	actualStatus, _ := status.FromError(err)

	assert.Equal(t, codes.InvalidArgument, actualStatus.Code())
	assert.Equal(t, errDescribeOfficeIDValidation, actualStatus.Message())
}

func Test_officeAPI_DescribeOfficeV1_Without_ID(t *testing.T) {
	t.Parallel()

	fixture := setUp(t)
	_, err := fixture.apiServer.DescribeOfficeV1(context.Background(), &bss_office_api.DescribeOfficeV1Request{})

	actualStatus, _ := status.FromError(err)

	assert.Equal(t, codes.InvalidArgument, actualStatus.Code())
	assert.Equal(t, errDescribeOfficeIDValidation, actualStatus.Message())
}
