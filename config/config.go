package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Host string
		Port string
	}
	Database struct {
		Host     string
		Port     int
		Name     string
		User     string
		Password string
		SSLMode  string
	}
	JWT struct {
		SecretKey            string
		TokenLifetimeMinutes int
	}
	Email struct {
		SMTPServer string
		SMTPPort   int
		Username   string
		Password   string
		From       string
	}
}

var AppConfig Config

func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading file, %s", err)
	}
	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
}
