package cmd

import (
	http "github.com/nivaldogmelo/web-api-tester/pkg/http"
)

func StartServer(port string) {
	http.Serve(port)
}
