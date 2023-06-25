package config

import (
	"log"

	models "github.com/mazeyqian/go-gin-gee/internal/pkg/models/sites"
	"github.com/spf13/viper"
)

var Config *Configuration

type Configuration struct {
	Server   ServerConfiguration
	Database DatabaseConfiguration
	Data     DataConfiguration
}

type DatabaseConfiguration struct {
	Driver       string
	Dbname       string
	Username     string
	Password     string
	Host         string
	Port         string
	MaxLifetime  int
	MaxOpenConns int
	MaxIdleConns int
}

type ServerConfiguration struct {
	Port   string
	Secret string
	Mode   string
}

type DataConfiguration struct {
	Sites []models.WebSite
}

// SetupDB initialize configuration
func Setup(configPath string, configType string) {
	var configuration *Configuration

	viper.SetConfigFile(configPath)
	// https://pkg.go.dev/github.com/spf13/viper@v1.13.0#SetConfigType
	viper.SetConfigType(configType) // "yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	Config = configuration
}

// GetConfig helps you to get configuration data
func GetConfig() *Configuration {
	return Config
}
