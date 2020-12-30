package sqlite

import (
	"database/sql"
	"encoding/json"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/nivaldogmelo/web-api-tester/internal/root"
)

func InitDB() error {
	database, err := sql.Open("sqlite3", "config/database.db")
	defer database.Close()
	if err != nil {
		log.Println("Error opening sqlite instance")
		return err
	}

	statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS requests (id INTEGER PRIMARY KEY, name TEXT, method TEXT, headers TEXT, body TEXT)")
	if err != nil {
		return err
	}

	_, err = statement.Exec()
	if err != nil {
		return err
	}

	return nil
}

func InsertRequest(request root.Request) error {
	database, err := sql.Open("sqlite3", "config/database.db")
	defer database.Close()
	if err != nil {
		log.Println("Error opening sqlite instance")
		return err
	}

	statement, err := database.Prepare("INSERT INTO requests (name, method, headers, body) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Println("Error inserting new request")
		return err
	}

	header, err := json.Marshal(request.Headers)
	if err != nil {
		log.Fatal(err)
	}

	_, err = statement.Exec(request.Name, request.Method, header, request.Body)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func GetAllRequests() ([]root.Request, error) {
	database, err := sql.Open("sqlite3", "config/database.db")
	defer database.Close()
	if err != nil {
		log.Println("Error opening sqlite instance")
		return nil, err
	}

	rows, err := database.Query("SELECT * FROM requests")
	if err != nil {
		log.Println("Error getting requests from database")
		return nil, err
	}

	var id int
	var name string
	var method string
	var headers string
	var body string
	var request []root.Request
	for rows.Next() {
		rows.Scan(&id, &name, &method, &headers, &body)

		var header []root.Header
		err = json.Unmarshal([]byte(headers), &header)

		temp := root.Request{
			Name:    name,
			Method:  method,
			Headers: header,
			Body:    body,
		}

		request = append(request, temp)
	}

	return request, nil
}
