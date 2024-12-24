package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Config struct {
	Env string `yaml:"env" env-default:"prod"`
}

type ENVConfig struct {
	DBHost        string `env:"DB_HOST"`
	DBPort        string `env:"DB_PORT"`
	DBUser        string `env:"DB_USER"`
	DBPassword    string `env:"DB_PASSWORD"`
	DBDbname      string `env:"DB_DBNAME"`
	AppPort       string `env:"APP_PORT"`
	AppEnv        string `env:"APP_ENV"`
	CacheHost     string `env:"CACHE_HOST"`
	CachePort     string `env:"CACHE_PORT"`
	JwtSalt       string `env:"JWT_Salt"`
	JwtSigningKey string `env:"JWT_SigningKey"`
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
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", &err)
	}
	return cfg
}

func LoadEnv() (ENVConfig, error) {
	err := godotenv.Load()

	if err != nil {
		return ENVConfig{}, fmt.Errorf("error loading .env file: %w", err)
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		return ENVConfig{}, fmt.Errorf("DB_HOST is not set")
	}
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		return ENVConfig{}, fmt.Errorf("DB_PORT is not set")
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		return ENVConfig{}, fmt.Errorf("DB_NAME is not set")

	}
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		return ENVConfig{}, fmt.Errorf("DB_USER is not set")
	}
	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		return ENVConfig{}, fmt.Errorf("DB_PASSWORD is not set")
	}
	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		return ENVConfig{}, fmt.Errorf("APP_PORT is not set")
	}
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		return ENVConfig{}, fmt.Errorf("APP_ENV is not set")
	}
	cacheHost := os.Getenv("CACHE_HOST")
	if dbHost == "" {
		return ENVConfig{}, fmt.Errorf("cache_HOST is not set")
	}
	cachePort := os.Getenv("CACHE_PORT")
	if dbHost == "" {
		return ENVConfig{}, fmt.Errorf("cache_PORT is not set")
	}
	jwtSalt := os.Getenv("JWT_Salt")
	if jwtSalt == "" {
		return ENVConfig{}, fmt.Errorf("JWT_Salt is not set")
	}
	jwtSigningKey := os.Getenv("JWT_SigningKey")
	if jwtSigningKey == "" {
		return ENVConfig{}, fmt.Errorf("JWT_SigningKey is not set")
	}

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		dbHost,
		dbPort,
		dbUser,
		dbName,
		dbPassword,
		appEnv,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return ENVConfig{}, fmt.Errorf("error connecting to database: %w", err)
	}
	if err := db.Ping(); err != nil {
		return ENVConfig{}, err
	}
	return ENVConfig{
		DBHost:        dbHost,
		DBPort:        dbPort,
		DBUser:        dbUser,
		DBPassword:    dbPassword,
		DBDbname:      dbName,
		AppPort:       appPort,
		AppEnv:        appEnv,
		CacheHost:     cacheHost,
		CachePort:     cachePort,
		JwtSalt:       jwtSalt,
		JwtSigningKey: jwtSigningKey,
	}, nil
}
