package requests

import (
	"time"

	"github.com/nivaldogmelo/web-api-tester/internal/root"
	vegeta "github.com/tsenart/vegeta/v12/lib"
)

func Execute(request root.Request) (vegeta.Metrics, error) {
	headers := transformHeaders(request.Headers)

	rate := vegeta.Rate{Freq: 1, Per: time.Second}
	duration := 4 * time.Second
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: request.Method,
		URL:    request.URL,
		Body:   []byte(request.Body),
		Header: headers,
	})
	attacker := vegeta.NewAttacker()

	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "Big Bang!") {
		metrics.Add(res)
	}
	metrics.Close()

	return metrics, nil
}

func transformHeaders(headers []root.Header) map[string][]string {
	finalHeaders := make(map[string][]string)

	for _, v := range headers {
		finalHeaders[v.Header] = []string{v.Content}
	}

	return finalHeaders
}
