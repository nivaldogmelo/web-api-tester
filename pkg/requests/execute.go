package requests

import (
	"bytes"
	"net/http"
	"strconv"
	"time"

	"github.com/nivaldogmelo/web-api-tester/internal/root"
	error_handler "github.com/nivaldogmelo/web-api-tester/pkg/error"
	"github.com/tcnksm/go-httpstat"
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

	var result httpstat.Result
	ctx := httpstat.WithHTTPStat(httpRequest.Context(), &result)
	httpRequest = httpRequest.WithContext(ctx)

	response, err := client.Do(httpRequest)
	if err != nil {
		if response == nil {
			end := time.Now()
			CountWebRequests(request.Name, strconv.Itoa(http.StatusBadGateway), result.Total(end))
			return http.StatusBadGateway, err
		} else {
			end := time.Now()
			CountWebRequests(request.Name, strconv.Itoa(response.StatusCode), result.Total(end))
			return response.StatusCode, err
		}
	}

	end := time.Now()
	CountWebRequests(request.Name, strconv.Itoa(response.StatusCode), result.Total(end))

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
