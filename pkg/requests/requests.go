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

func GetOne(name string) (root.Request, error) {
	var request root.Request

	request, err := sqlite.GetOneRequest(name)
	if err != nil {
		log.Println("Error getting request")
		return request, err
	}

	log.Println(request)

	return request, nil
}
