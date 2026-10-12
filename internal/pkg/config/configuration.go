package config

import (
	"flag"

	modelsLink "github.com/chengchuu/go-gin-gee/internal/pkg/models/link"
	modelsS "github.com/chengchuu/go-gin-gee/internal/pkg/models/sites"
	"github.com/chengchuu/go-gin-gee/pkg/logger"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var Config *Configuration

type Configuration struct {
	Server   ServerConfiguration
	Database DatabaseConfiguration
	Data     DataConfiguration
}

type ServerConfiguration struct {
	Port string
	Mode string
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
	EnableCORS          string
	WebhookID           string
	WebhookToken        string
	EnableWebhookAPI    string
	WebhookAPIKeys      []string
	CommonAPIKeys       []string
	LinkRedirectPageURL string
	KVAPIKeys           []string
	BaseURL             string
	AgentRecordsPath    string
	Sites               []modelsS.WebSite
	SpecialLinks        []modelsLink.SpecialLink
}

// SetupDB initialize configuration
func Setup() {
	var configuration *Configuration
	var err error

	// Flags
	flag.String("config-path", "data/config.json", "path of configuration")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	// BASE_URL is the only configuration environment override.
	if err = viper.BindEnv("BASE_URL"); err != nil {
		logger.Fatal("Unable to bind BASE_URL, %v", err)
	}
	// Default value
	viper.SetDefault("BASE_URL", "")
	viper.SetDefault("Data.EnableCORS", "off")
	viper.SetDefault("Data.EnableWebhookAPI", "off")

	// Configuration File
	configPath := viper.GetString("config-path")
	viper.SetConfigFile(configPath)
	viper.SetConfigType("json")
	// Read the configuration file
	if err = viper.ReadInConfig(); err != nil {
		logger.Println("No configuration file found, using default configuration")
	}
	err = viper.Unmarshal(&configuration)
	if err != nil {
		logger.Fatal("Unable to decode into struct, %v", err)
	}

	// A nonempty BASE_URL overrides the JSON value.
	baseURL := viper.GetString("BASE_URL")
	if baseURL != "" {
		configuration.Data.BaseURL = baseURL
	}

	// Fallback
	if configuration.Server.Port == "" {
		configuration.Server.Port = "3000"
	}
	if configuration.Server.Mode == "" {
		configuration.Server.Mode = "release"
	}

	Config = configuration
}

// GetConfig helps you to get configuration data
func GetConfig() *Configuration {
	if Config != nil && Config.Server.Mode == "debug" {
		logger.Info("Config.Server: %+v", Config.Server)
		// logger.Info("Config.Data.EnableWebhookAPI: %+v", Config.Data.EnableWebhookAPI)
	}
	return Config
}
