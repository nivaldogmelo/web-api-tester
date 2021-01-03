package sqlite

import (
	"database/sql"
	"encoding/json"
	"errors"

	_ "github.com/mattn/go-sqlite3"
	"github.com/nivaldogmelo/web-api-tester/internal/root"
	error_handler "github.com/nivaldogmelo/web-api-tester/pkg/error"
)

func InitDB() error {
	database, err := sql.Open("sqlite3", "config/database.db")
	defer database.Close()
	if err != nil {
		error_handler.Print(errors.New("Error opening database instance"))
		return err
	}

	statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS requests (id INTEGER PRIMARY KEY, name TEXT, method TEXT, headers TEXT, body TEXT, url TEXT)")
	if err != nil {
		error_handler.Print(errors.New("Error preparing database query"))
		return err
	}

	_, err = statement.Exec()
	if err != nil {
		error_handler.Print(errors.New("Error executing database query"))
		return err
	}

	return nil
}

func InsertRequest(request root.Request) error {
	database, err := sql.Open("sqlite3", "config/database.db")
	defer database.Close()
	if err != nil {
		error_handler.Print(errors.New("Error opening database instance"))
		return err
	}

	statement, err := database.Prepare("INSERT INTO requests (name, method, headers, body, url) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		error_handler.Print(errors.New("Error inserting new request in database"))
		return err
	}

	header, err := json.Marshal(request.Headers)
	if err != nil {
		error_handler.Print(errors.New("Error parsing JSON from headers"))
		return err
	}

	_, err = statement.Exec(request.Name, request.Method, header, request.Body, request.URL)
	if err != nil {
		error_handler.Print(errors.New("Error executing database query"))
		return err
	}

	return nil
}

func GetAllRequests() ([]root.Request, error) {
	database, err := sql.Open("sqlite3", "config/database.db")
	defer database.Close()
	if err != nil {
		error_handler.Print(errors.New("Error opening database instance"))
		return nil, err
	}

	rows, err := database.Query("SELECT * FROM requests")
	if err != nil {
		error_handler.Print(errors.New("Error getting requests from database"))
		return nil, err
	}

	var id int
	var name string
	var method string
	var headers string
	var body string
	var url string
	var request []root.Request

	for rows.Next() {
		rows.Scan(&id, &name, &method, &headers, &body, &url)

		var header []root.Header
		err = json.Unmarshal([]byte(headers), &header)

		temp := root.Request{
			Name:    name,
			Method:  method,
			Headers: header,
			Body:    body,
			URL:     url,
		}

		request = append(request, temp)
	}

	return request, nil
}

func GetOneRequest(id string) (root.Request, error) {
	var request root.Request

	database, err := sql.Open("sqlite3", "config/database.db")
	defer database.Close()
	if err != nil {
		error_handler.Print(errors.New("Error opening database instance"))
		return request, err
	}

	row := database.QueryRow("SELECT name, method, headers, body, url FROM requests WHERE id='" + id + "'")
	if err != nil {
		error_handler.Print(errors.New("Error getting request from database"))
		return request, err
	}

	var name string
	var method string
	var headers string
	var body string
	var url string

	err = row.Scan(&name, &method, &headers, &body, &url)
	if err != nil {
		return request, err
	}

	var header []root.Header
	err = json.Unmarshal([]byte(headers), &header)

	request = root.Request{
		Name:    name,
		Method:  method,
		Headers: header,
		Body:    body,
		URL:     url,
	}

	return request, nil
}
