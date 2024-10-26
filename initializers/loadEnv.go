package initializers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var ServerAdderss string

func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ServerAdderss = os.Getenv("SERVER_ADDRESS")
}
