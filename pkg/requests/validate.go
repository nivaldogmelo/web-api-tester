package requests

import (
	"errors"

	"github.com/nivaldogmelo/web-api-tester/internal/root"
	error_handler "github.com/nivaldogmelo/web-api-tester/pkg/error"
	"github.com/nivaldogmelo/web-api-tester/pkg/sqlite"
)

func checkIfDuplicated(newRequest root.Request) error {
	var reqNotFound = errors.New("No request found")

	requests, err := sqlite.GetRequestByField("name", newRequest.Name)
	if err != nil {
		if err != reqNotFound {
			error_handler.Print(errors.New("Error accessing requests database"))
			return nil
		}
	}

	for _, v := range requests {
		if v.Name == newRequest.Name {
			return errors.New("Request already exists")
		}
	}

	return nil
}
