package utils

import (
	"github.com/joho/godotenv"
	"os"
)

func GoDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		return ""
	}

	return os.Getenv(key)
}
