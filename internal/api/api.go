package api

import (
	"github.com/ozonmp/bss-office-api/internal/service/office"
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

type officeAPI struct {
	pb.UnimplementedBssOfficeApiServiceServer
	service office.IOfficeService
}

// NewOfficeAPI returns api of bss-office-api service
func NewOfficeAPI(s office.IOfficeService) pb.BssOfficeApiServiceServer {
	return &officeAPI{service: s}
}
