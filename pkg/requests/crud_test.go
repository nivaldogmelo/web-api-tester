package requests_test

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"
	"testing"

	"github.com/nivaldogmelo/web-api-tester/internal/root"
	"github.com/nivaldogmelo/web-api-tester/pkg/requests"
	"github.com/stretchr/testify/assert"
)

var exampleRequest = root.Request{
	Name:   "test",
	Method: "GET",
	Headers: []root.Header{
		{
			Header:  "Content-Type",
			Content: "application/json",
		},
	},
	Body: "username=lorem",
	URL:  "http://localhost:8080/",
}

func TestSave(t *testing.T) {
	initDatabaseForTests("test_database.db")

	newRequest := root.Request{
		Name:   "request",
		Method: "GET",
		Headers: []root.Header{
			{
				Header:  "Content-Type",
				Content: "application/json",
			},
		},
		Body: "username=lorem",
		URL:  "http://localhost:8080/",
	}

	err := requests.Save(newRequest)
	if err != nil {
		t.Errorf("Error saving request - %v", err)
	}

	err = os.Remove("test_database.db")
	if err != nil {
		t.Errorf("Error deleting created DB - %v", err)
	}
}

func TestGetAll(t *testing.T) {
	initDatabaseForTests("test_database.db")

	expectedResult := []root.Request{exampleRequest}

	actualResult, err := requests.GetAll()
	if err != nil {
		t.Errorf("Error getting requests - %v", err)
	}

	assert.Equal(t, expectedResult, actualResult)

	err = os.Remove("test_database.db")
	if err != nil {
		t.Errorf("Error deleting created DB - %v", err)
	}
}

func TestGetOne(t *testing.T) {
	initDatabaseForTests("test_database.db")

	actualResult, err := requests.GetOne("1")
	if err != nil {
		t.Errorf("Error getting request - %v", err)
	}

	assert.Equal(t, exampleRequest, actualResult)

	err = os.Remove("test_database.db")
	if err != nil {
		t.Errorf("Error deleting created DB - %v", err)
	}
}

func TestDeleteOne(t *testing.T) {
	initDatabaseForTests("test_database.db")

	err := requests.DeleteOne("1")
	if err != nil {
		t.Errorf("Error deleting request - %v", err)
	}

	err = os.Remove("test_database.db")
	if err != nil {
		t.Errorf("Error deleting created DB - %v", err)
	}
}

func initDatabaseForTests(dbFile string) {
	database, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatal(err)
	}

	defer database.Close()

	statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS requests (id INTEGER PRIMARY KEY, name TEXT, method TEXT, headers TEXT, body TEXT, url TEXT)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = statement.Exec()
	if err != nil {
		log.Fatal(err)
	}

	statement, err = database.Prepare("INSERT INTO requests (name, method, headers, body, url) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	header, err := json.Marshal(exampleRequest.Headers)
	if err != nil {
		log.Fatal(err)
	}

	_, err = statement.Exec(exampleRequest.Name, exampleRequest.Method, header, exampleRequest.Body, exampleRequest.URL)
	if err != nil {
		log.Fatal(err)
	}
}
