package main

import (
	"errors"
	"log"

	web "github.com/nivaldogmelo/web-api-tester/cmd"
	c "github.com/nivaldogmelo/web-api-tester/internal/config"
	error_handler "github.com/nivaldogmelo/web-api-tester/pkg/error"
	sqlite "github.com/nivaldogmelo/web-api-tester/pkg/sqlite"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")

	var config c.Config
	if err := viper.ReadInConfig(); err != nil {
		error_handler.Print(errors.New("Error reading config file, using default name"))
		log.Fatal(err)
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		error_handler.Print(errors.New("Error parsing config file using default name"))
		log.Fatal(err)
	}

	err = sqlite.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting to serve at port " + config.Server.Port + "...")
	web.StartServer(":" + config.Server.Port)
}
