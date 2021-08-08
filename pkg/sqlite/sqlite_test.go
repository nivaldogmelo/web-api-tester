package sqlite_test

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"
	"testing"

	"github.com/nivaldogmelo/web-api-tester/internal/root"
	error_handler "github.com/nivaldogmelo/web-api-tester/pkg/error"
	"github.com/nivaldogmelo/web-api-tester/pkg/sqlite"
	"github.com/stretchr/testify/assert"
)

var expectedRequest = root.Request{
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

type Table struct {
	CID          int
	Name         string
	Type         string
	NotNull      sql.NullString
	DefaultValue string
	PK           int
}

func TestInitDB(t *testing.T) {
	var expectedResult, actualResult Table

	initDatabaseForTests("expected_test_database.db")
	expectedDatabase, err := sql.Open("sqlite3", "expected_test_database.db")
	if err != nil {
		t.Errorf("Error getting expected DB - %v", err)
	}

	defer expectedDatabase.Close()

	expectedRows, _ := expectedDatabase.Query("PRAGMA table_info(requests)")

	err = sqlite.InitDB()
	if err != nil {
		t.Errorf("Error initializing DB - %v", err)
	}
	actualDatabase, err := sql.Open("sqlite3", "test_database.db")
	if err != nil {
		t.Errorf("Error getting expected DB - %v", err)
	}

	defer actualDatabase.Close()

	actualRows, _ := actualDatabase.Query("PRAGMA table_info(requests)")

	for actualRows.Next() {
		expectedRows.Next()

		expectedRows.Scan(&expectedResult.CID, &expectedResult.Name, &expectedResult.Type, &expectedResult.NotNull, &expectedResult.DefaultValue, &expectedResult.PK)
		actualRows.Scan(&actualResult.CID, &actualResult.Name, &actualResult.Type, &actualResult.NotNull, &actualResult.DefaultValue, &actualResult.PK)

		assert.Equal(t, expectedResult, actualResult)
	}

	err = os.Remove("test_database.db")
	if err != nil {
		t.Errorf("Error deleting test DB - %v", err)
	}
	err = os.Remove("expected_test_database.db")
	if err != nil {
		t.Errorf("Error deleting created DB - %v", err)
	}
}

func TestInsertRequests(t *testing.T) {
	err := sqlite.InitDB()
	if err != nil {
		t.Errorf("Error initializing request at DB - %v", err)
	}

	err = sqlite.InsertRequest(expectedRequest)
	if err != nil {
		t.Errorf("Error inserting request at DB - %v", err)
	}

	database, err := sql.Open("sqlite3", "test_database.db")
	if err != nil {
		t.Errorf("Error opening test DB - %v", err)
	}

	defer database.Close()

	actualRequest, err := sqlite.GetOneRequest("1")
	if err != nil {
		t.Errorf("Error getting request from DB - %v", err)
	}

	assert.Equal(t, expectedRequest, actualRequest)

	err = os.Remove("test_database.db")
	if err != nil {
		t.Errorf("Error deleting test DB - %v", err)
	}
}

func TestGetAllRequests(t *testing.T) {
	initDatabaseForTests("test_database.db")

	actualRequest, err := sqlite.GetAllRequests()
	if err != nil {
		t.Errorf("Error getting requests from DB - %v", err)
	}

	expectedResult := []root.Request{expectedRequest}

	assert.Equal(t, expectedResult, actualRequest)

	err = os.Remove("test_database.db")
	if err != nil {
		t.Errorf("Error deleting test DB - %v", err)
	}
}

func TestGetOneRequest(t *testing.T) {
	initDatabaseForTests("test_database.db")

	actualRequest, err := sqlite.GetOneRequest("1")
	if err != nil {
		t.Errorf("Error getting request from DB - %v", err)
	}

	assert.Equal(t, expectedRequest, actualRequest)

	err = os.Remove("test_database.db")
	if err != nil {
		t.Errorf("Error deleting test DB - %v", err)
	}
}

func TestGetRequestByField(t *testing.T) {
	initDatabaseForTests("test_database.db")

	actualRequest, err := sqlite.GetRequestByField("name", "teste")
	if err != nil {
		t.Errorf("Error getting request from DB - %v", err)
	}

	expectedResult := []root.Request{expectedRequest}

	assert.Equal(t, expectedResult, actualRequest)

	err = os.Remove("test_database.db")
	if err != nil {
		t.Errorf("Error deleting test DB - %v", err)
	}
}

func TestDeleteOneRequest(t *testing.T) {
	initDatabaseForTests("test_database.db")

	err := sqlite.DeleteOneRequest("1")
	if err != nil {
		t.Errorf("Error deleting request from DB - %v", err)
	}

	_, expectedResult := sqlite.GetOneRequest("1")

	assert.Equal(t, "sql: no rows in result set", expectedResult.Error())

	err = os.Remove("test_database.db")
	if err != nil {
		t.Errorf("Error deleting test DB - %v", err)
	}
}

func initDatabaseForTests(dbFile string) {
	_, err := os.Create(dbFile)
	if err != nil {
		log.Println(err)
	}

	database, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatal(err)
	}

	defer database.Close()

	statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS requests (id INTEGER PRIMARY KEY, name TEXT, method TEXT, headers TEXT, body TEXT, url TEXT)")
	if err != nil {
		log.Println(database)
		error_handler.Print(err)
		log.Fatal(err)
	}

	_, err = statement.Exec()
	if err != nil {
		error_handler.Print(err)
		log.Fatal(err)
	}

	statement, err = database.Prepare("INSERT INTO requests (name, method, headers, body, url) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	header, err := json.Marshal(expectedRequest.Headers)
	if err != nil {
		log.Fatal(err)
	}

	_, err = statement.Exec(expectedRequest.Name, expectedRequest.Method, header, expectedRequest.Body, expectedRequest.URL)
	if err != nil {
		log.Fatal(err)
	}
}
