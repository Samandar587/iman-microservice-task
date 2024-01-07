package config

import "os"

type Config struct {
	App         string
	Environment string
	LogLevel    string
	RPCPort     string
	HTTPPort    string

	DB struct {
		Host     string
		Port     string
		Database string
		User     string
		Password string
		Sslmode  string
	}
}

func NewDB() *Config {
	var config Config

	config.App = getEnv("APP", "iman-microservice-task")
	config.Environment = getEnv("ENVIRONMENT", "develop")
	config.LogLevel = getEnv("RPC_PORT", ":5005")
	config.HTTPPort = getEnv("HTTP_PORT", ":5005")

	config.DB.Host = getEnv("POSTGRES_HOST", "localhost")
	config.DB.Port = getEnv("POSTGRES_PORT", "5432")
	config.DB.Database = getEnv("POSTGRES_DATABASE", "blog_db")
	config.DB.User = getEnv("POSTGRES_USER", "postgres")
	config.DB.Password = getEnv("POSTGRES_PASSWORD", "postgres")
	config.DB.Sslmode = getEnv("POSTGRES_SSLMODE", "disable")
	return &config
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
