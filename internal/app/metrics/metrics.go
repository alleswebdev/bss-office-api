package metrics

import (
	"github.com/ozonmp/bss-office-api/internal/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var totalEventsProcessing prometheus.Gauge

// InitMetrics - инициализирует метрики
func InitMetrics(cfg config.Config) {
	totalEventsProcessing = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: cfg.Metrics.Namespace,
		Subsystem: cfg.Metrics.Subsystem,
		Name:      "events_processing_total",
		Help:      "Total number of the events in processing",
	})
}

// AddEventsTotal - увеличивает счетчик обрабатываемых событий
func AddEventsTotal(count float64) {
	if totalEventsProcessing == nil {
		return
	}

	totalEventsProcessing.Add(count)
}

// SubEventsTotal - уменьшает счетчик обрабатываемых событий
func SubEventsTotal(count float64) {
	if totalEventsProcessing == nil {
		return
	}

	totalEventsProcessing.Sub(count)
}
