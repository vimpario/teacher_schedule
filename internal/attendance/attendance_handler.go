package attendance

import (
	"encoding/json"
	"net/http"
	"strconv"
	"teacher_schedule/pkg/config"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
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
	var attendance Attendance
	err := json.NewDecoder(r.Body).Decode(&attendance)
	if err != nil{
		http.Error(w, "Неверный запрос", http.StatusInternalServerError)
		return
	}

	err = config.DB.Create(&attendance).Error
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(attendance)
}

func UpdateAttndanceHandler(w http.ResponseWriter, r *http.Request){
	id := chi.URLParam(r, "id")
	var attendance Attendance

	err := config.DB.First(&attendance, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Посещаемость не найдена", http.StatusNotFound)
		} else{
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	var updatetdData Attendance
	err = json.NewDecoder(r.Body).Decode(&updatetdData)
	if err != nil{
		http.Error(w, "Неверное тело запроса", http.StatusBadRequest)	
		return
	}

	err = config.DB.Model(&attendance).Updates(updatetdData).Error
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(attendance)
}

func AddBulkAttendancesHandler(w http.ResponseWriter, r *http.Request){
	var newAttendances []NewAttendance

	err := json.NewDecoder(r.Body).Decode(&newAttendances)
	if err != nil{
		http.Error(w, "Неверное тело запроса", http.StatusBadRequest)
		return
	}

	var attendances []Attendance

	for _, attendance := range newAttendances{
		newAttendance := Attendance{
			StudentID: attendance.StudentID,
			ScheduleID: attendance.ScheduleID,
			IsPresent: false,
		}
		if attendance.IsPresent != nil {
			newAttendance.IsPresent = *attendance.IsPresent
		}
		attendances = append(attendances, newAttendance)
	}

	err = config.DB.Create(&attendances).Error
	if err != nil {
		http.Error(w, "ошибка при добавлении расписания", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(attendances)
}

func UpdateBulkAttendancesHandler(w http.ResponseWriter, r *http.Request){
	var attendacesID [] uint
	err := json.NewDecoder(r.Body).Decode(&attendacesID)
	if err != nil {
		http.Error(w, "Неверное тело запроса", http.StatusBadRequest)
		return
	}
	if len(attendacesID) == 0{
		http.Error(w, "Пустое тело запроса", http.StatusBadRequest)
		return
	}
	
}

func DeleteAttendanceHandler(w http.ResponseWriter, r *http.Request){
	id := chi.URLParam(r, "id")
	err := config.DB.Delete(&Attendance{}, id).Error
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteBulkAttendacesHandler(w http.ResponseWriter, r *http.Request){
	var attandecesID [] uint
	err := json.NewDecoder(r.Body).Decode(&attandecesID)
	if err != nil{
		http.Error(w, "Неверное тело запроса", http.StatusBadRequest)
		return
	}
	if len(attandecesID) == 0{
		http.Error(w, "Пустое тело запроса", http.StatusBadRequest)
		return
	}

	result := config.DB.Where("attendanceid in ?", attandecesID).Delete(&Attendance{})
	if result.Error != nil{
		http.Error(w, "Ошибка удаления указаныых посещаемостей", http.StatusInternalServerError)
		return
	}
	response := map[string]interface{}{
		"message": "Посещаемости удалены успешно",
		"count": result.RowsAffected,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
