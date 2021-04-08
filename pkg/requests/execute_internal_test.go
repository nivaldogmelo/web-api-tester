package requests

import (
	"testing"

	"github.com/nivaldogmelo/web-api-tester/internal/root"
	"github.com/stretchr/testify/assert"
)

func TestTransformHeaders(t *testing.T) {
	exampleHeader := []root.Header{
		{
			Header:  "Content-Type",
			Content: "application/json",
		},
	}

	expectedHeaders := make(map[string][]string)
	expectedHeaders["Content-Type"] = []string{"application/json"}

	resultHeaders := transformHeaders(exampleHeader)

	assert.Equal(t, expectedHeaders, resultHeaders)
}
