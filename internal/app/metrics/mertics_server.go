package metrics

import (
	"fmt"
	"github.com/ozonmp/bss-office-api/internal/config"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

// CreateMetricsServer - создаёт http-endpoint с метриками
func CreateMetricsServer(cfg *config.Config) *http.Server {
	addr := fmt.Sprintf("%s:%d", cfg.Metrics.Host, cfg.Metrics.Port)

	mux := http.DefaultServeMux
	mux.Handle(cfg.Metrics.Path, promhttp.Handler())

	metricsServer := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	return metricsServer
}
