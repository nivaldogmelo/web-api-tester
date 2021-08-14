package requests

import (
	"bytes"
	"net/http"
	"strconv"
	"time"

	"github.com/nivaldogmelo/web-api-tester/internal/root"
	error_handler "github.com/nivaldogmelo/web-api-tester/pkg/error"
	vegeta "github.com/tsenart/vegeta/v12/lib"
)

func ExecuteRequest(request root.Request) (int, error) {
	headers := transformHeaders(request.Headers)

	client := &http.Client{}

	httpRequest, err := http.NewRequest(request.Method, request.URL, bytes.NewBufferString(request.Body))
	if err != nil {
		error_handler.Print(err)
		return http.StatusBadRequest, err
	}

	httpRequest.Header = headers

	response, err := client.Do(httpRequest)
	if err != nil {
		error_handler.Print(err)
		return response.StatusCode, err
	}

	CountWebRequests(request.Name, strconv.Itoa(response.StatusCode))

	return response.StatusCode, nil
}

func Load(request root.Request, frequency int, testDuration time.Duration) (vegeta.Metrics, error) {
	headers := transformHeaders(request.Headers)

	rate := vegeta.Rate{Freq: frequency, Per: time.Second}
	duration := testDuration * time.Second
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
