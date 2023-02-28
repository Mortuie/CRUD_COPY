package utils

import (
	"log"

	"github.com/caarlos0/env/v7"
	"github.com/joho/godotenv"
	"github.com/mortuie/CRUD_COPY/models"
)

func GetEnvVars() models.EnvVariables {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	cfg := models.EnvVariables{}

	err = env.Parse(&cfg)

	if err != nil {
		log.Fatal(err)
	}

	return cfg
}
