package requests

import (
	"log"

	"github.com/nivaldogmelo/web-api-tester/internal/root"
	"github.com/nivaldogmelo/web-api-tester/pkg/sqlite"
)

func Save(request root.Request) error {
	err := sqlite.InsertRequest(request)
	if err != nil {
		log.Println("Error saving request")
		return err
	}

	return nil
}

func GetAll() ([]root.Request, error) {
	requests, err := sqlite.GetAllRequests()
	if err != nil {
		log.Println("Error getting requests")
		return nil, err
	}

	return requests, nil
}
