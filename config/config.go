package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func (c *Config) readConfig() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}

	// API Config for localhost:8888
	api := os.Getenv("API_URL")
	c.ApiConfig = ApiConfig{Url: api}

	// DB Config start here
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASS")

	// Create the MySQL connection string (DSN)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Open a MySQL connection
	c.DbConfig = DbConfig{dsn}

	// c.FilePathConfig = FilePathConfig{FilePath: os.Getenv("FILE_PATH")}
}

func NewConfig() Config {
	cfg := Config{}
	cfg.readConfig()
	return cfg
}

func InitConfig() Config {
	cfg := Config{}
	cfg.readConfig()
	return cfg
}
