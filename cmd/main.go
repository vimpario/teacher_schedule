package main

import (
	"log"
	"net/http"
	"teacher_schedule/internal/attendance"
	"teacher_schedule/internal/schedule"
	"teacher_schedule/internal/subjects"
	"teacher_schedule/internal/users"
	"teacher_schedule/pkg/config"

	"teacher_schedule/pkg/httputil"
)

func main() {
    config.LoadConfig()
	config.InitDB()



	err := config.DB.AutoMigrate(&users.User{}, &schedule.Schedule{}, &subjects.Subjects{}, &attendance.Attendance{})
	if err != nil{
		log.Fatalf("ОШибка миграции моделей: %v", err)
	}

    router := httputil.NewRouter()
    log.Println("Сервер запущен на порту 8080....")
    http.ListenAndServe(":8080", router)
}
