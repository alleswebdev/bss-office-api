package metrics

import (
	"github.com/ozonmp/bss-office-api/internal/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var totalEventsProcessing prometheus.Gauge
var totalEventsProcessed prometheus.Counter

// InitMetrics - инициализирует метрики
func InitMetrics(cfg config.Config) {
	totalEventsProcessing = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: cfg.Metrics.Namespace,
		Subsystem: cfg.Metrics.Subsystem,
		Name:      "events_processing_total",
		Help:      "Total number of the events in processing",
	})

	totalEventsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: cfg.Metrics.Namespace,
		Subsystem: cfg.Metrics.Subsystem,
		Name:      "events_processed_total",
		Help:      "Total number of the processed events",
	})
}

// AddEventsProcessingTotal - увеличивает счетчик обрабатываемых событий
func AddEventsProcessingTotal(count float64) {
	if totalEventsProcessing == nil {
		return
	}

	totalEventsProcessing.Add(count)
}

// SubEventsProcessingTotal - уменьшает счетчик обрабатываемых событий
func SubEventsProcessingTotal(count float64) {
	if totalEventsProcessing == nil {
		return
	}

	totalEventsProcessing.Sub(count)
}

// AddEventsProcessedTotal - увеличивает счетчик обработанных событий
func AddEventsProcessedTotal(count float64) {
	if totalEventsProcessing == nil {
		return
	}

	totalEventsProcessed.Add(count)
}
