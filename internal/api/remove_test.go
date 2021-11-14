package api

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/jmoiron/sqlx"
	"github.com/ozonmp/bss-office-api/internal/model"
	bss_office_api "github.com/ozonmp/bss-office-api/pkg/bss-office-api"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

const errRemoveOfficeIDValidation = "invalid RemoveOfficeV1Request.OfficeId: value must be greater than 0"

func Test_officeAPI_RemoveOfficeV1(t *testing.T) {
	t.Parallel()

	fixture := setUp(t)

	testOfficeID := uint64(1)

	fixture.dbMock.ExpectBegin()
	fixture.dbMock.ExpectCommit()

	fixture.officeRepo.EXPECT().RemoveOffice(gomock.Any(), testOfficeID, gomock.Any()).DoAndReturn(func(ctx context.Context, officeID uint64, tx *sqlx.Tx) (bool, error) {
		return true, nil
	})

	fixture.eventRepo.EXPECT().Add(gomock.Any(), gomock.Eq(&model.OfficeEvent{
		OfficeID: testOfficeID,
		Type:     model.Removed,
		Status:   model.Deferred,
		Payload: model.OfficePayload{
			ID:      testOfficeID,
			Removed: true,
		},
	}))

	res, err := fixture.apiServer.RemoveOfficeV1(context.Background(),
		&bss_office_api.RemoveOfficeV1Request{OfficeId: testOfficeID})

	assert.True(t, res.GetFound())
	assert.NoError(t, err)
}

func Test_officeAPI_RemoveOfficeV1_Repo_Err(t *testing.T) {
	t.Parallel()

	fixture := setUp(t)

	testOfficeID := uint64(1)

	errTest := errors.New("test officeRepo err")

	fixture.dbMock.ExpectBegin()

	fixture.officeRepo.EXPECT().RemoveOffice(gomock.Any(), testOfficeID, gomock.Any()).
		Return(false, errTest)

	res, err := fixture.apiServer.RemoveOfficeV1(context.Background(),
		&bss_office_api.RemoveOfficeV1Request{OfficeId: testOfficeID})

	fixture.dbMock.ExpectCommit()

	actualStatus, _ := status.FromError(err)

	assert.Equal(t, codes.Internal, actualStatus.Code())
	assert.Error(t, errTest, actualStatus.Err())
	assert.False(t, res.GetFound())
}

func Test_officeAPI_RemoveOfficeV1_With_Zero_ID(t *testing.T) {
	t.Parallel()

	fixture := setUp(t)
	_, err := fixture.apiServer.RemoveOfficeV1(context.Background(),
		&bss_office_api.RemoveOfficeV1Request{OfficeId: 0})

	actualStatus, _ := status.FromError(err)

	assert.Equal(t, codes.InvalidArgument, actualStatus.Code())
	assert.Equal(t, errRemoveOfficeIDValidation, actualStatus.Message())
}

func Test_officeAPI_RemoveOfficeV1_Without_ID(t *testing.T) {
	t.Parallel()

	fixture := setUp(t)
	_, err := fixture.apiServer.RemoveOfficeV1(context.Background(), &bss_office_api.RemoveOfficeV1Request{})

	actualStatus, _ := status.FromError(err)

	assert.Equal(t, codes.InvalidArgument, actualStatus.Code())
	assert.Equal(t, errRemoveOfficeIDValidation, actualStatus.Message())
}
