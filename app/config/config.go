package config

import (
	"fmt"

	"github.com/3cognito/library/app/utils"
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

const (
	Dev = "dev"
)

var Configs *Config

type Database struct {
	Host     string `env:"HOST,required"`
	Port     int    `env:"PORT,required"`
	User     string `env:"USER,required"`
	Password string `env:"PASSWORD,required"`
	Name     string `env:"NAME,required"`
}

type Config struct {
	ENV                       string   `env:"ENV" envDefault:"dev"`
	Port                      string   `env:"PORT,required"`
	DB                        Database `env:"" envPrefix:"DB_"`
	AccessTokenExpiryDuration string   `env:"ACCESS_TOKEN_EXPIRY_DURATION,required"` //should be an integer in hours
	AppJWTSecret              string   `env:"APP_JWT_SECRET,required"`
}

func Load() {
	if loadErr := godotenv.Load(); loadErr != nil {
		fmt.Println("Error loading config file: ", loadErr)
		panic(loadErr)
	}

	config := Config{}

	if parseErr := env.Parse(&config); parseErr != nil {
		fmt.Println("Error parsing config: ", parseErr)
		panic(parseErr)
	}

	Configs = &config

	utils.ParseAccessTokenExpiryTime(config.AccessTokenExpiryDuration) //panics if the access token expiry duration is not a valid number
}
