package config

import (
	
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(){
	dsn := "host=" + GetEnv("DB_HOST") +
        " user=" + GetEnv("DB_USER") +
        " password=" + GetEnv("DB_PASSWORD") +
        " dbname=" + GetEnv("DB_NAME") +
        " port=" + GetEnv("DB_PORT") +
        " sslmode=disable"

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil{
		log.Fatalf("Ошибка подключения к бд: %v", err)
	}

	log.Println("Подключение к БД успешно")
}