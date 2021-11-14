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

const testLimit = uint64(1)
const testOffset = uint64(1)

const errListOfficeValidation = "invalid ListOfficesV1Request.Limit: value must be inside range (0, 100)"

func Test_officeAPI_ListOfficeV1(t *testing.T) {
	t.Parallel()

	fixture := setUp(t)

	testOfficeID := uint64(1)
	testOfficeName := "test name"

	fixture.officeRepo.EXPECT().ListOffices(gomock.Any(), gomock.Eq(testLimit), gomock.Eq(testOffset)).
		Return([]*model.Office{{
			ID:   testOfficeID,
			Name: testOfficeName,
		}}, nil)

	res, err := fixture.apiServer.ListOfficesV1(context.Background(),
		&bss_office_api.ListOfficesV1Request{Limit: testLimit, Offset: testOffset})

	assert.Equal(t, res.GetItems()[0].GetId(), testOfficeID)
	assert.Equal(t, res.GetItems()[0].GetName(), testOfficeName)
	assert.NoError(t, err)
}

func Test_officeAPI_ListOfficeV1_Repo_Err(t *testing.T) {
	t.Parallel()

	fixture := setUp(t)

	errTest := errors.New("test officeRepo err")
	fixture.officeRepo.EXPECT().ListOffices(gomock.Any(), gomock.Eq(testLimit), gomock.Eq(testOffset)).
		DoAndReturn(func(ctx context.Context, limit uint64, offset uint64) ([]*model.Office, error) {
			return nil, errTest
		})

	res, err := fixture.apiServer.ListOfficesV1(context.Background(),
		&bss_office_api.ListOfficesV1Request{Limit: testLimit, Offset: testOffset})

	actualStatus, _ := status.FromError(err)

	assert.Equal(t, codes.Internal, actualStatus.Code())
	assert.Error(t, errTest, actualStatus.Err())
	assert.Nil(t, res.GetItems())
}

func Test_officeAPI_ListOfficeV1_With_Zero_Limit(t *testing.T) {
	t.Parallel()

	fixture := setUp(t)
	_, err := fixture.apiServer.ListOfficesV1(context.Background(),
		&bss_office_api.ListOfficesV1Request{Limit: 0, Offset: testOffset})

	actualStatus, _ := status.FromError(err)

	assert.Equal(t, codes.InvalidArgument, actualStatus.Code())
	assert.Equal(t, errListOfficeValidation, actualStatus.Message())
}

func Test_officeAPI_ListOfficeV1_With_Large_Limit(t *testing.T) {
	t.Parallel()

	fixture := setUp(t)
	_, err := fixture.apiServer.ListOfficesV1(context.Background(),
		&bss_office_api.ListOfficesV1Request{Limit: 101, Offset: testOffset})

	actualStatus, _ := status.FromError(err)

	assert.Equal(t, codes.InvalidArgument, actualStatus.Code())
	assert.Equal(t, errListOfficeValidation, actualStatus.Message())
}
