package requests

import (
	"errors"

	"github.com/nivaldogmelo/web-api-tester/internal/root"
	error_handler "github.com/nivaldogmelo/web-api-tester/pkg/error"
	"github.com/nivaldogmelo/web-api-tester/pkg/sqlite"
)

func Save(request root.Request) error {
	err := sqlite.InsertRequest(request)
	if err != nil {
		error_handler.Print(errors.New("Error saving request"))
		return err
	}

	return nil
}

func GetAll() ([]root.Request, error) {
	requests, err := sqlite.GetAllRequests()
	if err != nil {
		error_handler.Print(errors.New("Error gettings requests"))
		return nil, err
	}

	return requests, nil
}

func GetOne(name string) (root.Request, error) {
	var request root.Request

	request, err := sqlite.GetOneRequest(name)
	if err != nil {
		error_handler.Print(errors.New("Error getting requests"))
		return request, err
	}

	return request, nil
}
