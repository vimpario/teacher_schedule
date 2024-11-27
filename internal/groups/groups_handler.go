package groups

import (
	"encoding/json"
	"net/http"
	"teacher_schedule/pkg/config"
)

func GetGroupHandler(w http.ResponseWriter, r *http.Request) {
	var groups []Group
	result := config.DB.Find(&groups)
	if result.Error != nil{
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(groups)
	
}

func CreateGroupHandler(w http.ResponseWriter, r *http.Request) {
	var group Group
	err:=json.NewDecoder(r.Body).Decode(&group)
	if err != nil{
		http.Error(w, "Неверный ввод", http.StatusBadRequest)
		return
	}

	result := config.DB.Create(&group)
	if result.Error != nil{
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(group)
}