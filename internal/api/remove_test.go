package api

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	bss_office_api "github.com/ozonmp/bss-office-api/pkg/bss-office-api"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

const errRemoveOfficeIdValidation = "invalid RemoveOfficeV1Request.OfficeId: value must be greater than 0"

func Test_officeAPI_RemoveOfficeV1(t *testing.T) {
	t.Parallel()

	fixture := setUp(t)

	testOfficeID := uint64(1)

	fixture.repo.EXPECT().RemoveOffice(gomock.Any(), testOfficeID).DoAndReturn(func(ctx context.Context, officeID uint64) (bool, error) {
		return true, nil
	})

	res, err := fixture.apiServer.RemoveOfficeV1(context.Background(),
		&bss_office_api.RemoveOfficeV1Request{OfficeId: testOfficeID})

	assert.True(t, res.GetFound())
	assert.NoError(t, err)
}

func Test_officeAPI_RemoveOfficeV1_Repo_Err(t *testing.T) {
	t.Parallel()

	fixture := setUp(t)

	testOfficeID := uint64(1)

	errTest := errors.New("test repo err")
	fixture.repo.EXPECT().RemoveOffice(gomock.Any(), testOfficeID).
		DoAndReturn(func(ctx context.Context, officeID uint64) (bool, error) {
			return false, errTest
		})

	res, err := fixture.apiServer.RemoveOfficeV1(context.Background(),
		&bss_office_api.RemoveOfficeV1Request{OfficeId: testOfficeID})

	assert.False(t, res.GetFound())
	assert.Error(t, err, errTest)
}

func Test_officeAPI_RemoveOfficeV1_With_Zero_ID(t *testing.T) {
	t.Parallel()

	fixture := setUp(t)
	_, err := fixture.apiServer.RemoveOfficeV1(context.Background(),
		&bss_office_api.RemoveOfficeV1Request{OfficeId: 0})

	actualStatus, _ := status.FromError(err)

	assert.Equal(t, codes.InvalidArgument, actualStatus.Code())
	assert.Equal(t, errRemoveOfficeIdValidation, actualStatus.Message())
}

func Test_officeAPI_RemoveOfficeV1_Without_ID(t *testing.T) {
	t.Parallel()

	fixture := setUp(t)
	_, err := fixture.apiServer.RemoveOfficeV1(context.Background(), &bss_office_api.RemoveOfficeV1Request{})

	actualStatus, _ := status.FromError(err)

	assert.Equal(t, codes.InvalidArgument, actualStatus.Code())
	assert.Equal(t, errRemoveOfficeIdValidation, actualStatus.Message())
}
