package api

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/ozonmp/bss-office-api/internal/mocks"
	"github.com/ozonmp/bss-office-api/internal/model"
	bss_office_api "github.com/ozonmp/bss-office-api/pkg/bss-office-api"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

const errCreateNameValidation = "invalid CreateOfficeV1Request.Name: value length must be between 2 and 100 runes, inclusive"

type APIFixture struct {
	repo      *mocks.MockRepo
	ctrl      *gomock.Controller
	apiServer bss_office_api.BssOfficeApiServiceServer
}

func setUp(t *testing.T) APIFixture {
	var fixture APIFixture

	fixture.ctrl = gomock.NewController(t)
	fixture.repo = mocks.NewMockRepo(fixture.ctrl)
	fixture.apiServer = NewOfficeAPI(fixture.repo)

	return fixture
}

func Test_officeAPI_CreateOfficeV1(t *testing.T) {
	t.Parallel()
	fixture := setUp(t)

	testName := "Office 5"
	testID := uint64(1)

	fixture.repo.EXPECT().CreateOffice(gomock.Any(), model.Office{Name: testName}).DoAndReturn(func(ctx context.Context, office model.Office) (uint64, error) {
		return testID, nil
	})

	res, err := fixture.apiServer.CreateOfficeV1(context.Background(), &bss_office_api.CreateOfficeV1Request{Name: testName})

	assert.Equal(t, testID, res.GetOfficeId())
	assert.NoError(t, err)
}

func Test_officeAPI_CreateOfficeV1_Repo_Err(t *testing.T) {
	t.Parallel()

	fixture := setUp(t)

	testName := "test office"
	errTest := errors.New("test repo err")

	fixture.repo.EXPECT().CreateOffice(gomock.Any(), model.Office{Name: testName}).
		DoAndReturn(func(ctx context.Context, office model.Office) (*model.Office, error) {
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

	_, err := fixture.apiServer.CreateOfficeV1(context.Background(), &bss_office_api.CreateOfficeV1Request{Name: "Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Aenean commodo ligula eget dolor. Aenean mass"})

	actualStatus, _ := status.FromError(err)

	assert.Equal(t, codes.InvalidArgument, actualStatus.Code())
	assert.Equal(t, errCreateNameValidation, actualStatus.Message())
}
