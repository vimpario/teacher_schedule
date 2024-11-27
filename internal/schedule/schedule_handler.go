package schedule

import (
	"encoding/json"
	"net/http"
	"teacher_schedule/pkg/config"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func CreateScheduleHandler(w http.ResponseWriter, r *http.Request){
	var schedule Schedule
	err := json.NewDecoder(r.Body).Decode(&schedule)
	if err != nil{
		http.Error(w, "Неверный запрос", http.StatusBadRequest)
		return
	}

	err  = config.DB.Create(&schedule).Error
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(schedule)
}

func GetAllSchedulesHandler(w http.ResponseWriter, r *http.Request){
	var schedules []Schedule
	err := config.DB.Find(&schedules).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(schedules)
}

func GetScheduleByIdHandler(w http.ResponseWriter, r *http.Request){
	id := chi.URLParam(r, "id")
	var schedule Schedule

	err := config.DB.First(&schedule, id).Error
	if err != nil{
		if err == gorm.ErrRecordNotFound{
			http.Error(w, "Расписание не найдено", http.StatusNotFound)
		}else{
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	json.NewEncoder(w).Encode(schedule)
}

func UpdateScheduleHandler(w http.ResponseWriter, r *http.Request){
	id := chi.URLParam(r, "id")
	var schedule Schedule

	err := config.DB.First(&schedule, id).Error
	if err != nil{
		if err == gorm.ErrRecordNotFound{
			http.Error(w, "Расписание не найдено", http.StatusNotFound)
		}else{
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	var updatedData Schedule
	err = json.NewDecoder(r.Body).Decode(&updatedData)
	if err != nil{
		http.Error(w, "Неверное тело запроса", http.StatusBadRequest)
		return
	}

	err = config.DB.Model(&schedule).Updates(updatedData).Error
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(schedule)
}

func DeleteScheduleHandler(w http.ResponseWriter, r *http.Request){
	id := chi.URLParam(r, "id")
	err := config.DB.Delete(&Schedule{}, id).Error
	if err!=nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}