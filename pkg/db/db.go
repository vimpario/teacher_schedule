package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(connectionString string) *gorm.DB{
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil{
		log.Fatalf("Ошибка подключения БДЖ %v", err)
	}

	return db
}