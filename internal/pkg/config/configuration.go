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

type ServerConfiguration struct {
	Port   string
	Secret string
	Mode   string
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

type DataConfiguration struct {
	Sites []models.WebSite
	// WECOM_ROBOT_CHECK
	WeComRobotCheck string
}

// SetupDB initialize configuration
func Setup(configPath string, configType string) {
	var configuration *Configuration

	// Automatically read environment variables that match
	viper.AutomaticEnv()
	// Set the default value
	viper.SetDefault("WECOM_ROBOT_CHECK", "")
	// Config File
	viper.SetConfigFile(configPath)
	// https://pkg.go.dev/github.com/spf13/viper@v1.13.0#SetConfigType
	viper.SetConfigType(configType)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	// Supply the environment variables.
	weComRobotCheck := viper.GetString("WECOM_ROBOT_CHECK")
	log.Println("configuration.Data.WeComRobotCheck:", configuration.Data.WeComRobotCheck)
	log.Println("weComRobotCheck:", weComRobotCheck)
	if weComRobotCheck != "" {
		configuration.Data.WeComRobotCheck = weComRobotCheck
	}
	Config = configuration
}

// GetConfig helps you to get configuration data
func GetConfig() *Configuration {
	return Config
}
