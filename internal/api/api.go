package api

import (
	"context"
	"github.com/ozonmp/bss-office-api/internal/model"
	pb "github.com/ozonmp/bss-office-api/pkg/bss-office-api"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	totalOfficeNotFound = promauto.NewCounter(prometheus.CounterOpts{
		Name: "bss_office_api_office_not_found_total",
		Help: "Total number of offices that were not found",
	})
)

type OfficeService interface {
	RemoveOffice(ctx context.Context, officeID uint64) (bool, error)
	DescribeOffice(ctx context.Context, officeID uint64) (*model.Office, error)
	ListOffices(ctx context.Context, limit uint64, offset uint64) ([]*model.Office, error)
	CreateOffice(ctx context.Context, office model.Office) (uint64, error)
}

type officeAPI struct {
	pb.UnimplementedBssOfficeApiServiceServer
	service OfficeService
}

// NewOfficeAPI returns api of bss-office-api service
func NewOfficeAPI(s OfficeService) pb.BssOfficeApiServiceServer {
	return &officeAPI{service: s}
}
