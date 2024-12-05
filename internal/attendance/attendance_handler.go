package attendance

import (
	"encoding/json"
	"net/http"
	"strconv"
	"teacher_schedule/pkg/config"

	"github.com/go-chi/chi/v5"
)

func GetAllAttendancesHandler(w http.ResponseWriter, r *http.Request){
	var attendances []Attendance
	err := config.DB.Find(&attendances).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(attendances)
}

func GetAttendanceByStudentHandler(w http.ResponseWriter, r *http.Request){
	studentIDParam := chi.URLParam(r, "studentId")
	studentID, err := strconv.Atoi(studentIDParam)

	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var attendances []Attendance
	err = config.DB.Where("studentid = ?", studentID).Find(&attendances).Error
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(attendances) == 0{
		http.Error(w, "Для данного студента нет посещаемости", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(attendances)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetAttendanceByScheduleHandler(w http.ResponseWriter, r *http.Request){
	scheduleIDParam := chi.URLParam(r, "scheduleId")
	scheduleID, err := strconv.Atoi(scheduleIDParam)

	if err != nil{
		http.Error(w, "Неверный ID расписания", http.StatusBadRequest)
		return
	}

	var attendances []Attendance
	err = config.DB.Where("scheduleid = ?", scheduleID).Find(&attendances).Error
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(attendances) == 0{
		http.Error(w, "Для данного расписания нет посещемости", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(attendances)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func AddAttendanceHandler(w http.ResponseWriter, r *http.Request){

}

func UpdateAttndanceHandler(w http.ResponseWriter, r *http.Request){

}

func AddBulkAttendancesHandler(w http.ResponseWriter, r *http.Request){

}

func UpdateBulkAttendancesHandler(w http.ResponseWriter, r *http.Request){

}

func DeleteAttendanceHandler(w http.ResponseWriter, r *http.Request){

}

func DeleteBulkAttendacesHandler(w http.ResponseWriter, r *http.Request){

}
