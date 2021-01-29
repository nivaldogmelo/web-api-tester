package requests

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"
	"testing"

	"github.com/nivaldogmelo/web-api-tester/internal/root"
	"github.com/stretchr/testify/assert"
)

var request = root.Request{
	Name:   "teste",
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

func TestCheckIfDuplicated(t *testing.T) {
	initDatabaseForTests("database.db")

	result := checkIfDuplicated(request)

	assert.Equal(t, "Request already exists", result.Error())

	err := os.Remove("database.db")
	if err != nil {
		t.Errorf("Error deleting created DB - %v", err)
	}
}

func initDatabaseForTests(dbFile string) {
	database, err := sql.Open("sqlite3", dbFile)
	defer database.Close()
	if err != nil {
		log.Fatal(err)
	}

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

	header, err := json.Marshal(request.Headers)
	if err != nil {
		log.Fatal(err)
	}

	_, err = statement.Exec(request.Name, request.Method, header, request.Body, request.URL)
	if err != nil {
		log.Fatal(err)
	}
}
