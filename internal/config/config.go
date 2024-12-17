package config

import (
	"database/sql"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type Config struct {
	Env string `yaml:"env" env-default:"prod"`
}

type ENVConfig struct {
	dbHost     string `env:"DB_HOST"`
	dbPort     int    `env:"DB_PORT"`
	dbUser     string `env:"DB_USER"`
	dbPassword string `env:"DB_PASSWORD"`
	dbDbname   string `env:"DB_DBNAME"`
	appPort    int    `env:"APP_PORT"`
	appEnv     string `env:"APP_ENV"`
}

func NewConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatal("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", &err)
	}
	return cfg
}

func StorageConfig() (ENVConfig, error) {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		log.Fatal("DB_HOST is not set")
	}
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		log.Fatal("DB_PORT is not set")
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		log.Fatal("DB_NAME is not set")

	}
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		log.Fatal("DB_USER is not set")
	}
	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		log.Fatal("DB_PASSWORD is not set")
	}
	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		log.Fatal("APP_PORT is not set")
	}
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		log.Fatal("APP_ENV is not set")
	}
	return ENVConfig{
		dbHost:     dbHost,
		dbPort:     dbPort,
		dbUser:     dbUser,
		dbPassword: dbPassword,
		dbDbname:   dbName,
		appPort:    appPort,
		appEnv:     appEnv,
	}, nil
}

func ConnectFile(*sql.DB, error) {
	connStr := fmt.Sprintf(
		`host=%s port=%s name=%s user=%s password=%s app_port=%s app_env=%s`,
		dbHost, dbPort, dbName, dbUser, dbPassword, appPort, appEnv)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return db, nil
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
}
