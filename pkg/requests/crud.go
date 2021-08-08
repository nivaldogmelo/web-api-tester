package requests

import (
	"errors"

	"github.com/nivaldogmelo/web-api-tester/internal/root"
	error_handler "github.com/nivaldogmelo/web-api-tester/pkg/error"
	"github.com/nivaldogmelo/web-api-tester/pkg/sqlite"
)

func Save(request root.Request) error {
	err := checkIfDuplicated(request)
	if err != nil {
		return err
	}

	err = sqlite.InsertRequest(request)
	if err != nil {
		error_handler.Print(errors.New("error saving request"))
		return err
	}

	return nil
}

func GetAll() ([]root.Request, error) {
	requests, err := sqlite.GetAllRequests()
	if err != nil {
		error_handler.Print(errors.New("error getting requests"))
		return nil, err
	}

	return requests, nil
}

func GetOne(id string) (root.Request, error) {
	var request root.Request

	request, err := sqlite.GetOneRequest(id)
	if err != nil {
		if err.Error() != "sql: no rows in result set" {
			error_handler.Print(errors.New("error getting request"))
		}
		return request, err
	}

	return request, nil
}

func DeleteOne(id string) error {
	err := sqlite.DeleteOneRequest(id)
	if err != nil {
		error_handler.Print(errors.New("error deleting request"))
		return err
	}

	return nil
}
