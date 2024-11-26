package main

import (
	"fmt"
	"teacher_schedule/internal/schedule"
	"teacher_schedule/internal/users"
	"teacher_schedule/pkg/config"
	"teacher_schedule/pkg/db"
)

func main() {
    config.LoadConfig()

    connectionString := "host=" + config.GetEnv("DB_HOST") +
        " user=" + config.GetEnv("DB_USER") +
        " password=" + config.GetEnv("DB_PASSWORD") +
        " dbname=" + config.GetEnv("DB_NAME") +
        " port=" + config.GetEnv("DB_PORT") +
        " sslmode=disable"
    database := db.Connect(connectionString)

    database.AutoMigrate(&users.User{}, &schedule.Schedule{})
	
}