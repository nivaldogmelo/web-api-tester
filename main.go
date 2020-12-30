package main

import (
	"log"

	web "github.com/nivaldogmelo/web-api-tester/cmd"
	sqlite "github.com/nivaldogmelo/web-api-tester/pkg/sqlite"
)

func main() {
	log.Println("Starting to serve at port 3000...")

	err := sqlite.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	web.StartServer(":3000")
}
