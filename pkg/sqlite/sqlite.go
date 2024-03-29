package sqlite

import (
	"database/sql"
	"encoding/json"
	"errors"
	"flag"

	_ "github.com/mattn/go-sqlite3"
	c "github.com/nivaldogmelo/web-api-tester/internal/config"
	"github.com/nivaldogmelo/web-api-tester/internal/root"
	error_handler "github.com/nivaldogmelo/web-api-tester/pkg/error"
	"github.com/spf13/viper"
)

func getDB() (string, error) {
	if flag.Lookup("test.v") != nil {
		return "test_database.db", nil
	}

	viper.SetConfigName("config")
	viper.AddConfigPath("config/")
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")

	var config c.Config
	if err := viper.ReadInConfig(); err != nil {
		error_handler.Print(errors.New("error reading config file, using default name"))
		return "database.db", err
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		error_handler.Print(errors.New("error parsing config file using default name"))
		return "database.db", err
	}

	return config.Database.Filename, nil
}

func OpenDB() (*sql.DB, error) {
	dbFile, err := getDB()
	if err != nil {
		error_handler.Print(errors.New("error getting database"))
		return nil, err
	}

	database, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		error_handler.Print(errors.New("error opening database instance"))
		return nil, err
	}

	return database, nil
}

func InitDB() error {
	database, err := OpenDB()
	if err != nil {
		error_handler.Print(errors.New("error at OpenDB()"))
	}

	defer database.Close()

	statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS requests (id INTEGER PRIMARY KEY, name TEXT, method TEXT, headers TEXT, body TEXT, url TEXT)")
	if err != nil {
		error_handler.Print(errors.New("error preparing database query"))
		return err
	}

	_, err = statement.Exec()
	if err != nil {
		error_handler.Print(errors.New("error executing database query"))
		return err
	}

	return nil
}

func InsertRequest(request root.Request) error {
	database, err := OpenDB()
	if err != nil {
		error_handler.Print(errors.New("error at OpenDB()"))
	}

	defer database.Close()

	statement, err := database.Prepare("INSERT INTO requests (name, method, headers, body, url) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		error_handler.Print(errors.New("error inserting new request in database"))
		return err
	}

	header, err := json.Marshal(request.Headers)
	if err != nil {
		error_handler.Print(errors.New("error parsing JSON from headers"))
		return err
	}

	_, err = statement.Exec(request.Name, request.Method, header, request.Body, request.URL)
	if err != nil {
		error_handler.Print(errors.New("error executing database query"))
		return err
	}

	return nil
}

func GetAllRequests() ([]root.Request, error) {
	database, err := OpenDB()
	if err != nil {
		error_handler.Print(errors.New("error at OpenDB()"))
	}

	defer database.Close()

	rows, err := database.Query("SELECT * FROM requests")
	if err != nil {
		error_handler.Print(errors.New("error getting requests from database"))
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
		if err != nil {
			error_handler.Print(errors.New("error building request headers"))
		}

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

	database, err := OpenDB()
	if err != nil {
		error_handler.Print(errors.New("error at OpenDB()"))
	}

	defer database.Close()

	row := database.QueryRow("SELECT name, method, headers, body, url FROM requests WHERE id='" + id + "'")
	if err != nil {
		error_handler.Print(errors.New("error getting request from database"))
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
	if err != nil {
		error_handler.Print(errors.New("error building request headers"))
	}

	request = root.Request{
		Name:    name,
		Method:  method,
		Headers: header,
		Body:    body,
		URL:     url,
	}

	return request, nil
}

func GetRequestByField(field string, value string) ([]root.Request, error) {
	database, err := OpenDB()
	if err != nil {
		error_handler.Print(errors.New("error at OpenDB()"))
	}

	defer database.Close()

	rows, err := database.Query("SELECT name, method, headers, body, url FROM requests WHERE " + field + "='" + value + "'")
	if err != nil {
		error_handler.Print(errors.New("error getting requests from database"))
		return nil, err
	}

	var name string
	var method string
	var headers string
	var body string
	var url string
	var request []root.Request

	for rows.Next() {
		rows.Scan(&name, &method, &headers, &body, &url)

		var header []root.Header
		err = json.Unmarshal([]byte(headers), &header)
		if err != nil {
			error_handler.Print(errors.New("error building request headers"))
		}

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

func DeleteOneRequest(id string) error {
	database, err := OpenDB()
	if err != nil {
		error_handler.Print(errors.New("error at OpenDB()"))
	}

	defer database.Close()

	statement, err := database.Prepare("DELETE FROM requests WHERE id=?")
	if err != nil {
		error_handler.Print(errors.New("error preparing delete query"))
		return err
	}

	result, err := statement.Exec(id)
	if err != nil {
		error_handler.Print(errors.New("error deleting request from database"))
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		err = errors.New("no request found")
		error_handler.Print(err)
		return err
	}

	return nil
}
