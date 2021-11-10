package api

import (
	"context"
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/jmoiron/sqlx"
	"github.com/ozonmp/bss-office-api/internal/mocks"
	"github.com/ozonmp/bss-office-api/internal/model"
	"github.com/ozonmp/bss-office-api/internal/service/office"
	bss_office_api "github.com/ozonmp/bss-office-api/pkg/bss-office-api"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

const errCreateNameValidation = "invalid CreateOfficeV1Request.Name: value length must be between 2 and 100 runes, inclusive"

type APIFixture struct {
	officeRepo    *mocks.MockOfficeRepo
	eventRepo     *mocks.MockEventRepo
	officeService office.OfficeService
	ctrl          *gomock.Controller
	apiServer     bss_office_api.BssOfficeApiServiceServer
	db            *sql.DB
	dbMock        sqlmock.Sqlmock
}

func setUp(t *testing.T) APIFixture {
	var fixture APIFixture

	fixture.ctrl = gomock.NewController(t)
	fixture.officeRepo = mocks.NewMockOfficeRepo(fixture.ctrl)
	fixture.eventRepo = mocks.NewMockEventRepo(fixture.ctrl)
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	fixture.db = db
	fixture.dbMock = mock

	fixture.officeService = office.NewOfficeService(fixture.officeRepo, fixture.eventRepo, sqlx.NewDb(db, "sqlmock"))
	fixture.apiServer = NewOfficeAPI(fixture.officeService)

	return fixture
}

func (f *APIFixture) tearDown() {
	f.db.Close()
}

func Test_officeAPI_CreateOfficeV1(t *testing.T) {
	t.Parallel()
	fixture := setUp(t)
	defer fixture.tearDown()

	testName := "Office 5"
	testID := uint64(1)

	fixture.dbMock.ExpectBegin()
	fixture.dbMock.ExpectCommit()

	fixture.eventRepo.EXPECT().Add(gomock.Any(), gomock.Eq(&model.OfficeEvent{
		OfficeID: testID,
		Type:     model.Created,
		Status:   model.Deferred,
		Payload: model.OfficePayload{
			ID:          testID,
			Name:        testName,
			Description: "",
			Removed:     false,
		},
	}))

	fixture.officeRepo.EXPECT().CreateOffice(gomock.Any(), model.Office{Name: testName}, gomock.Any()).DoAndReturn(func(ctx context.Context, office model.Office, tx *sqlx.Tx) (uint64, error) {
		return testID, nil
	})

	res, err := fixture.apiServer.CreateOfficeV1(context.Background(), &bss_office_api.CreateOfficeV1Request{Name: testName})

	assert.Equal(t, testID, res.GetOfficeId())
	assert.NoError(t, err)
}

func Test_officeAPI_CreateOfficeV1_Repo_Err(t *testing.T) {
	t.Parallel()

	fixture := setUp(t)
	defer fixture.tearDown()

	testName := "test office"
	errTest := errors.New("test officeRepo err")

	fixture.dbMock.ExpectBegin()
	fixture.dbMock.ExpectCommit()

	fixture.officeRepo.EXPECT().CreateOffice(gomock.Any(), model.Office{Name: testName}, gomock.Any()).
		DoAndReturn(func(ctx context.Context, office model.Office, tx *sqlx.Tx) (*model.Office, error) {
			return nil, errTest
		})

	_, err := fixture.apiServer.CreateOfficeV1(context.Background(),
		&bss_office_api.CreateOfficeV1Request{Name: testName})

	actualStatus, _ := status.FromError(err)

	assert.Equal(t, codes.Internal, actualStatus.Code())
	assert.Error(t, errTest, actualStatus.Err())
}

func Test_officeAPI_CreateOfficeV1_Error_Validation_Empty_Name(t *testing.T) {
	t.Parallel()

	fixture := setUp(t)
	defer fixture.tearDown()

	_, err := fixture.apiServer.CreateOfficeV1(context.Background(), &bss_office_api.CreateOfficeV1Request{Name: ""})

	actualStatus, _ := status.FromError(err)

	assert.Equal(t, codes.InvalidArgument, actualStatus.Code())
	assert.Equal(t, errCreateNameValidation, actualStatus.Message())
}

func Test_officeAPI_CreateOfficeV1_Error_Validation_Name_Min_Len(t *testing.T) {
	t.Parallel()
	fixture := setUp(t)

	_, err := fixture.apiServer.CreateOfficeV1(context.Background(), &bss_office_api.CreateOfficeV1Request{Name: "a"})

	actualStatus, _ := status.FromError(err)

	assert.Equal(t, codes.InvalidArgument, actualStatus.Code())
	assert.Equal(t, errCreateNameValidation, actualStatus.Message())
}

func Test_officeAPI_CreateOfficeV1_Error_Validation_Name_Max_Len(t *testing.T) {
	t.Parallel()
	fixture := setUp(t)
	defer fixture.tearDown()

	_, err := fixture.apiServer.CreateOfficeV1(context.Background(), &bss_office_api.CreateOfficeV1Request{Name: "Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Aenean commodo ligula eget dolor. Aenean mass"})

	actualStatus, _ := status.FromError(err)

	assert.Equal(t, codes.InvalidArgument, actualStatus.Code())
	assert.Equal(t, errCreateNameValidation, actualStatus.Message())
}
