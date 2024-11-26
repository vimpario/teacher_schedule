package users

import (
	"encoding/json"
	"net/http"
	"teacher_schedule/pkg/config"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var users []User
	result := config.DB.Find(&users)
	if result.Error != nil{
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil{
		http.Error(w, "Неверный ввод", http.StatusBadRequest)
		return
	}

	result := config.DB.Create(&user)
	if result.Error != nil{
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}