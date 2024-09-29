package configs

import (
	"log"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type EnvConfig struct {
	WORKDIR string `env:"WORKDIR,required"`
}

func loadDotEnv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	return nil
}

func InitConfig() (*EnvConfig, error) {

	// This command will return an error when deployed as a container, because a .env file does not exist
	err := loadDotEnv()
	if err != nil {
		log.Printf("Error loading .env file : %s", err)
	}

	// This command will run in a container, as the values are passed as envirnoment via the docker-compose.yaml file.
	cfg := EnvConfig{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil

}
