package api

import (
	"github.com/ozonmp/bss-office-api/internal/repo"
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
	repo repo.Repo
}

// NewOfficeAPI returns api of bss-office-api service
func NewOfficeAPI(r repo.Repo) pb.BssOfficeApiServiceServer {
	return &officeAPI{repo: r}
}
