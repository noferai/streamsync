package config

import (
	"github.com/spf13/viper"
	"log"
)

var Config *Configuration

type ServerConfiguration struct {
	Address string
	Port    string
}

type DatabaseConfiguration struct {
	Driver   string
	Dbname   string
	Username string
	Password string
	Host     string
	Port     string
	LogMode  bool
}

type Configuration struct {
	Server   ServerConfiguration
	Database DatabaseConfiguration
}

// Init configuration
func Setup() {
	var configuration *Configuration

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	Config = configuration
}

func GetConfig() *Configuration {
	return Config
}
