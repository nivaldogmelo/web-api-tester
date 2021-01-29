package http

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/nivaldogmelo/web-api-tester/internal/root"
)

func TestHomeHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(homeHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `"Hello World"`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestSaveHandler(t *testing.T) {
	initDatabaseForTests("database.db")

	body := strings.NewReader(`{
    "Name": "newTest",
    "Method": "GET",
    "Headers": [
        {
            "Header": "Content-Type",
            "Content": "application/json"
        },
        {
            "Header": "User"
        }
    ],
    "Body": "username=jason",
    "URL": "https://localhost:8080"
}`)

	req, err := http.NewRequest("POST", "/", body)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(saveHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	// Check the response body is what we expect.
	expected := `"Saved with Success"`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

	err = os.Remove("database.db")
	if err != nil {
		t.Errorf("Error deleting created DB - %v", err)
	}
}

func TestDeleteOneHandler(t *testing.T) {
	initDatabaseForTests("database.db")

	req, err := http.NewRequest("DELETE", "/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/{id}", deleteOneHandler)

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `"Request deleted with success"`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

	err = os.Remove("database.db")
	if err != nil {
		t.Errorf("Error deleting created DB - %v", err)
	}
}

func TestGetOneHandler(t *testing.T) {
	initDatabaseForTests("database.db")

	req, err := http.NewRequest("GET", "/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/{id}", getOneHandler)

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"Name":"test","Method":"GET","Headers":[{"Header":"Content-Type","Content":"application/json"}],"Body":"username=lorem","URL":"http://localhost:8080/"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

	err = os.Remove("database.db")
	if err != nil {
		t.Errorf("Error deleting created DB - %v", err)
	}
}

func TestGetAllHandler(t *testing.T) {
	initDatabaseForTests("database.db")

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getAllHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `[{"Name":"test","Method":"GET","Headers":[{"Header":"Content-Type","Content":"application/json"}],"Body":"username=lorem","URL":"http://localhost:8080/"}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

	err = os.Remove("database.db")
	if err != nil {
		t.Errorf("Error deleting created DB - %v", err)
	}
}

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

	header, err := json.Marshal(exampleRequest.Headers)
	if err != nil {
		log.Fatal(err)
	}

	_, err = statement.Exec(exampleRequest.Name, exampleRequest.Method, header, exampleRequest.Body, exampleRequest.URL)
	if err != nil {
		log.Fatal(err)
	}
}
