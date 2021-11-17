package api

import (
	"context"
	"github.com/ozonmp/bss-office-api/internal/model"
	pb "github.com/ozonmp/bss-office-api/pkg/bss-office-api"
)

type officeService interface {
	RemoveOffice(ctx context.Context, officeID uint64) (bool, error)
	DescribeOffice(ctx context.Context, officeID uint64) (*model.Office, error)
	ListOffices(ctx context.Context, limit uint64, offset uint64) ([]*model.Office, error)
	CreateOffice(ctx context.Context, office model.Office) (uint64, error)
	UpdateOffice(ctx context.Context, officeID uint64, office model.Office) (bool, error)
	UpdateOfficeName(ctx context.Context, officeID uint64, name string) (bool, error)
	UpdateOfficeDescription(ctx context.Context, officeID uint64, description string) (bool, error)
}

type officeAPI struct {
	pb.UnimplementedBssOfficeApiServiceServer
	service officeService
}

// NewOfficeAPI returns api of bss-office-api service
func NewOfficeAPI(s officeService) pb.BssOfficeApiServiceServer {
	return &officeAPI{service: s}
}
