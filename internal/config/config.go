package config

import (
	"encoding/json"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type LogParams struct {
	LogDirectory     string `json:"logDirectory"`
	LogInfo          string `json:"logInfo"`
	LogError         string `json:"logError"`
	LogWarn          string `json:"logWarn"`
	LogDebug         string `json:"logDebug"`
	MaxSizeMegabytes int    `json:"maxSizeMegabytes"`
	MaxBackups       int    `json:"maxBackups"`
	MaxAge           int    `json:"maxAge"`
	Compress         bool   `json:"compress"`
	LocalTime        bool   `json:"localTime"`
}

type FullConfig struct {
	DB          DBConfig  `json:"database"`
	LogParams   LogParams `json:"logParams"`
	DB_URL      string    `json:"-"`
	JWT_SECRET  string    `json:"-"`
	SERVER_PORT string    `json:"-"`
}

var AppConfig FullConfig

func LoadConfig() {
	_ = godotenv.Load()

	// Load from .env first
	AppConfig.JWT_SECRET = getEnv("JWT_SECRET", "1234567890")
	AppConfig.SERVER_PORT = getEnv("SERVER_PORT", "8181")
	AppConfig.DB_URL = getEnv("DB_URL", "")

	// Then override with configs.json values if present
	file, err := os.Open("internal/config/configs.json")
	if err != nil {
		log.Fatalf("❌ Failed to open configs.json: %v", err)
	}
	defer file.Close()

	tempConfig := struct {
		DB        DBConfig  `json:"database"`
		LogParams LogParams `json:"logParams"`
	}{}

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tempConfig); err != nil {
		log.Fatalf("❌ Failed to decode configs.json: %v", err)
	}

	AppConfig.DB = tempConfig.DB
	AppConfig.LogParams = tempConfig.LogParams
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
