package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"regexp"
)

const projectDirName = "gin-mongo-api"

func loadEnv() {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + `/.env`)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

}

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	URL      string
	Host     string
	Port     string
	Username string
	Password string
	Name     string
}

func GetConfig() *Config {
	loadEnv()
	return &Config{
		DB: &DBConfig{
			URL:      os.Getenv("DB_URL"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_DATABASE"),
		},
	}
}
