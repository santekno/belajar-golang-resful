package environment

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Load() (string, error) {
	fileEnv := ".env"
	if os.Getenv("environment") == "development" {
		fileEnv = "../.env"
	}

	err := godotenv.Load(fileEnv)
	if err != nil {
		log.Fatalf("error loading .env file")
	}

	return fileEnv, err
}
