package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	createCounter prometheus.Counter
	updateCounter prometheus.Counter
	removeCounter prometheus.Counter
)

func RegisterMetrics() {
	createCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "ocp_certificate_api_create_count_total",
		Help: "The total create certificate",
	})

	updateCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "ocp_certificate_api_update_count_total",
		Help: "The total update certificate",
	})

	removeCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "ocp_certificate_api_remove_count_total",
		Help: "The total remove certificate",
	})
}

func CreateCounterInc() {
	createCounter.Inc()
}

func UpdateCounterInc() {
	updateCounter.Inc()
}

func RemoveCounterInc() {
	removeCounter.Inc()
}
