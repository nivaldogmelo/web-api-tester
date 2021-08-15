package requests

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	WebRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "web_requests_total",
			Help: "How many registered web requests were made, partitioned by name and result",
		},
		[]string{"name", "code"},
	)

	WebRequestsLatency = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:       "web_requests_latency",
			Help:       "Latency of registered web requests",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		},
		[]string{"name", "code"},
	)
)

func CountWebRequests(requestName string, requestResult string, latency time.Duration) {
	WebRequests.WithLabelValues(requestName, requestResult).Inc()
	WebRequestsLatency.WithLabelValues(requestName, requestResult).Observe(latency.Seconds())
}
