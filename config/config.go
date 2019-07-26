package config

import (
	"fmt"
	"log"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

// Configuration contains static info required to run the apps
// It contains DB info
type Configuration struct {
	Address               string `env:"ADDRESS" envDefault:":8080"`
	JwtSecret             string `env:"JWT_SECRET,required"`
	DatabaseConnectionURL string `env:"CONNECTION_URL,required"`
}

// NewConfig will read the config data from given .env file
func NewConfig(files ...string) *Configuration {
	err := godotenv.Load(files...)

	if err != nil {
		log.Printf("No .env file could be found %q\n", files)
	}

	cfg := Configuration{}
	err = env.Parse(&cfg)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	return &cfg
}
