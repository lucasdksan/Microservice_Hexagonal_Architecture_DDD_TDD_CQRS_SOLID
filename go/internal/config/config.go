package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	Bank_connection_string = ""
	Port                   = 0
	Secret_key             []byte
)

func LoadingEnv() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("API_PORT"))

	if err != nil {
		Port = 8080
	}

	Bank_connection_string = os.Getenv("DATABASE_URL")

	Secret_key = []byte(os.Getenv("SECRET_KEY"))
}
