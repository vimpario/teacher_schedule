package main

import (
	"log"
	"net/http"
	"teacher_schedule/pkg/config"
	
	"teacher_schedule/pkg/httputil"
)

func main() {
    config.LoadConfig()

    router := httputil.NewRouter()
    log.Println("Сервер запущен на порту 8080....")
    http.ListenAndServe(":8080", router)
}
