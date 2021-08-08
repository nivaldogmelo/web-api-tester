package requests

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	WebRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "web_requests_total",
			Help: "How many registered web requests were made, partitioned by name and result",
		},
		[]string{"name", "result"},
	)
)

func CountWebRequests(requestName string, requestResult string) {
	WebRequests.WithLabelValues(requestName, requestResult).Inc()
}
