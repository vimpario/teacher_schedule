package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil{
		log.Println(".env Не обнаружен")
	}
}

func GetEnv(key string) string{
	return os.Getenv(key)
}