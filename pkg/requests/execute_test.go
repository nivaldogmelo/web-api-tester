package requests_test

import (
	"testing"

	"github.com/nivaldogmelo/web-api-tester/internal/root"
	"github.com/nivaldogmelo/web-api-tester/pkg/requests"
)

func TestExecute(t *testing.T) {

	metrics, err := requests.Execute(exampleRequest)
	if err != nil {
		t.Errorf("Error executing request - %v ", err)
	}

	err = root.TestStructType(metrics, "Metrics")
	if err != nil {
		t.Errorf("%s", err)
	}
}
