package metrics

import (
	"github.com/ozonmp/bss-office-api/internal/config"
	"github.com/ozonmp/bss-office-api/internal/model"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var totalOfficeNotFound prometheus.Counter
var totalCud *prometheus.CounterVec

// InitMetrics - инициализирует метрики
func InitMetrics(cfg *config.Config) {
	totalOfficeNotFound = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: cfg.Metrics.Namespace,
		Subsystem: cfg.Metrics.Subsystem,
		Name:      "not_found_total",
		Help:      "Total number of offices that were not found",
	})

	totalCud = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: cfg.Metrics.Namespace,
		Subsystem: cfg.Metrics.Subsystem,
		Name:      "cud_total",
		Help:      "Total of the CUD events",
	}, []string{"type"})
}

// IncTotalNotFound - увеличивает значение счетчика ошибок отсутствия объекта
func IncTotalNotFound() {
	if totalOfficeNotFound == nil {
		return
	}

	totalOfficeNotFound.Inc()
}

// IncTotalCud - увеличивает значение счетчика событий
func IncTotalCud(eventType model.EventType) {
	if totalCud == nil {
		return
	}

	totalCud.WithLabelValues(eventType.String()).Inc()
}
