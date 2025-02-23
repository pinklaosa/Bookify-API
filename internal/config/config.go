package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App      AppConfig
	Database DatabaseConfig
	Key KeyConfig
}

type AppConfig struct {
	Port int
}

type KeyConfig struct {
	SecretKey string
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBname   string
	Sslmode  string
}

var AppConfigInstance *Config

func LoadConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./internal/config")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Failed to decode config: %v", err)
	}

	AppConfigInstance = &config
	return &config
}
