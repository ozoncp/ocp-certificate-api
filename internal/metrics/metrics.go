package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Metrics interface {
	CreateCounterInc()
	UpdateCounterInc()
	RemoveCounterInc()
	MultiCreateCounterInc()
}

type metrics struct {
	createCounter      prometheus.Counter
	updateCounter      prometheus.Counter
	removeCounter      prometheus.Counter
	multiCreateCounter prometheus.Counter
}

func NewMetrics() *metrics {
	return &metrics{
		createCounter: promauto.NewCounter(prometheus.CounterOpts{
			Name: "ocp_certificate_api_create_count_total",
			Help: "The total create certificate",
		}),

		updateCounter: promauto.NewCounter(prometheus.CounterOpts{
			Name: "ocp_certificate_api_update_count_total",
			Help: "The total update certificate",
		}),

		removeCounter: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "ocp_certificate_api_remove_count_total",
			Help: "The total remove certificate",
		}),

		multiCreateCounter: promauto.NewCounter(prometheus.CounterOpts{
			Name: "ocp_certificate_api_multi_create_count_total",
			Help: "The total multi create certificate",
		}),
	}
}

func (m *metrics) CreateCounterInc() {
	m.createCounter.Inc()
}

func (m *metrics) UpdateCounterInc() {
	m.updateCounter.Inc()
}

func (m *metrics) RemoveCounterInc() {
	m.removeCounter.Inc()
}

func (m *metrics) MultiCreateCounterInc() {
	m.multiCreateCounter.Inc()
}
