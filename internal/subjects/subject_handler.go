package subjects

import (
	"encoding/json"
	"net/http"
	"teacher_schedule/pkg/config"
)

func GetSubjectHandler(w http.ResponseWriter, r *http.Request){
	var subjects []Subjects
	result := config.DB.Find(&subjects)
	if result.Error != nil{
		http.Error(w, result.Error.Error(),http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(subjects)
}

func CreateSunbjectHandler(w http.ResponseWriter, r *http.Request){
	var subject Subjects
	err := json.NewDecoder(r.Body).Decode(&subject)
	if err != nil{
		http.Error(w, "Неверный ввод", http.StatusBadRequest)
		return
	}

	result := config.DB.Create(&subject)
	if result.Error != nil{
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(subject)
}