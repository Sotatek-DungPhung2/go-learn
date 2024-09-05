package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var Port string

func ConfigEnv() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	Port = os.Getenv("PORT")

}
