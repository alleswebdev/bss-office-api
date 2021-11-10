package api

import (
	"context"
	"database/sql"
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

const errUpdateOfficeIDValidation = "invalid UpdateOfficeV1Request.OfficeId: value must be greater than 0"

var testOffice = model.Office{
	ID:          uint64(1),
	Name:        "test",
	Description: "test",
}

func Test_officeAPI_UpdateOfficeV1(t *testing.T) {
	t.Parallel()

	fixture := setUp(t)

	fixture.dbMock.ExpectBegin()
	fixture.dbMock.ExpectCommit()

	fixture.officeRepo.EXPECT().UpdateOffice(gomock.Any(), testOffice.ID, testOffice, gomock.Any()).
		DoAndReturn(func(ctx context.Context, officeID uint64, office model.Office, tx *sqlx.Tx) (bool, error) {
			return true, nil
		})

	fixture.eventRepo.EXPECT().Add(gomock.Any(), gomock.Eq(&model.OfficeEvent{
		OfficeID: testOffice.ID,
		Type:     model.Updated,
		Status:   model.Deferred,
		Payload: model.OfficePayload{
			ID:          testOffice.ID,
			Name:        testOffice.Name,
			Description: testOffice.Description,
		},
	}))

	res, err := fixture.apiServer.UpdateOfficeV1(context.Background(),
		&bss_office_api.UpdateOfficeV1Request{
			OfficeId:    testOffice.ID,
			Name:        testOffice.Name,
			Description: testOffice.Description,
		})

	assert.True(t, res.GetStatus())
	assert.NoError(t, err)
}

func Test_officeAPI_UpdateOfficeV1_Repo_Err(t *testing.T) {
	t.Parallel()

	fixture := setUp(t)

	testOfficeID := uint64(1)

	errTest := errors.New("test officeRepo err")

	fixture.dbMock.ExpectBegin()
	fixture.dbMock.ExpectCommit()

	fixture.officeRepo.EXPECT().UpdateOffice(gomock.Any(), testOfficeID, testOffice, gomock.Any()).
		DoAndReturn(func(ctx context.Context, officeID uint64, office model.Office, tx *sqlx.Tx) (bool, error) {
			return false, errTest
		})

	res, err := fixture.apiServer.UpdateOfficeV1(context.Background(),
		&bss_office_api.UpdateOfficeV1Request{
			OfficeId:    testOfficeID,
			Name:        testOffice.Name,
			Description: testOffice.Description,
		})

	actualStatus, _ := status.FromError(err)

	assert.Equal(t, codes.Internal, actualStatus.Code())
	assert.Error(t, errTest, actualStatus.Err())
	assert.False(t, res.GetStatus())
}

func Test_officeAPI_UpdateOfficeV1_Not_Found_Err(t *testing.T) {
	t.Parallel()

	fixture := setUp(t)

	testOfficeID := uint64(1)

	fixture.dbMock.ExpectBegin()
	fixture.dbMock.ExpectCommit()

	fixture.officeRepo.EXPECT().UpdateOffice(gomock.Any(), testOfficeID, testOffice, gomock.Any()).
		DoAndReturn(func(ctx context.Context, officeID uint64, office model.Office, tx *sqlx.Tx) (bool, error) {
			return false, sql.ErrNoRows
		})

	res, err := fixture.apiServer.UpdateOfficeV1(context.Background(),
		&bss_office_api.UpdateOfficeV1Request{
			OfficeId:    testOfficeID,
			Name:        testOffice.Name,
			Description: testOffice.Description,
		})

	actualStatus, _ := status.FromError(err)

	assert.Equal(t, codes.NotFound, actualStatus.Code())
	assert.Error(t, sql.ErrNoRows, actualStatus.Err())
	assert.False(t, res.GetStatus())
}

func Test_officeAPI_UpdateOfficeV1_With_Zero_ID(t *testing.T) {
	t.Parallel()

	fixture := setUp(t)
	_, err := fixture.apiServer.UpdateOfficeV1(context.Background(),
		&bss_office_api.UpdateOfficeV1Request{OfficeId: 0})

	actualStatus, _ := status.FromError(err)

	assert.Equal(t, codes.InvalidArgument, actualStatus.Code())
	assert.Equal(t, errUpdateOfficeIDValidation, actualStatus.Message())
}

func Test_officeAPI_UpdateOfficeV1_Without_ID(t *testing.T) {
	t.Parallel()

	fixture := setUp(t)
	_, err := fixture.apiServer.UpdateOfficeV1(context.Background(), &bss_office_api.UpdateOfficeV1Request{})

	actualStatus, _ := status.FromError(err)

	assert.Equal(t, codes.InvalidArgument, actualStatus.Code())
	assert.Equal(t, errUpdateOfficeIDValidation, actualStatus.Message())
}
