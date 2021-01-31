package sqlite

import (
	"testing"

	c "github.com/nivaldogmelo/web-api-tester/internal/config"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestGetDB(t *testing.T) {
	viper.SetConfigName("config")
	viper.AddConfigPath("../../config/")
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")

	var config c.Config
	if err := viper.ReadInConfig(); err != nil {
		t.Errorf("Error reading config file - %v", err)
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		t.Errorf("Error parsing config file - %v", err)
	}

	dbFilename, err := getDB()
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, config.Database.Filename, dbFilename)
}
