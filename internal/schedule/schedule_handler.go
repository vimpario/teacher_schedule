package schedule

import (
	"encoding/json"
	"net/http"
	"strconv"
	"teacher_schedule/pkg/config"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func CreateScheduleHandler(w http.ResponseWriter, r *http.Request) {
	var schedule Schedule
	err := json.NewDecoder(r.Body).Decode(&schedule)
	if err != nil {
		http.Error(w, "Неверный запрос", http.StatusBadRequest)
		return
	}

	err = config.DB.Create(&schedule).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(schedule)
}

func GetAllSchedulesHandler(w http.ResponseWriter, r *http.Request) {
	var schedules []Schedule
	err := config.DB.Find(&schedules).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(schedules)
}

func GetScheduleByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var schedule Schedule

	err := config.DB.First(&schedule, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Расписание не найдено", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	json.NewEncoder(w).Encode(schedule)
}

func GetScheduleByTeacherHandler(w http.ResponseWriter, r *http.Request) {
	teacherIDParam := chi.URLParam(r, "teacherId")
	teacherID, err := strconv.Atoi(teacherIDParam)

	if err != nil {
		http.Error(w, "Нверный ID преподавтеля", http.StatusBadRequest)
		return
	}

	var schedules []Schedule
	err = config.DB.Where("teacherid = ?", teacherID).Find(&schedules).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(schedules) == 0 {
		http.Error(w, "Для данного расписания нет расписания", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(schedules)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetScheduleFilteredHandler(w http.ResponseWriter, r *http.Request) {
	dayId := r.URL.Query().Get("dayId")
	groupId := r.URL.Query().Get("groupId")
	subjectId := r.URL.Query().Get("subjectId")
	isOccupied := r.URL.Query().Get("isOccupied")
	teacherId := r.URL.Query().Get("teacherId")
	sort := r.URL.Query().Get("sort")

	var schedules []Schedule
	query := config.DB.Model(&Schedule{})

	// query = ScheduleFilter()

	if dayId != "" {
		query = query.Where("dayid = ?", dayId)
	}
	if groupId != "" {
		query = query.Where("groupid = ?", groupId)
	}
	if subjectId != "" {
		query = query.Where("subjectid = ?", subjectId)
	}
	if isOccupied != "" {
		query = query.Where("isoccupied = ?", isOccupied)
	}
	if sort != "" {
		query = query.Order(sort)
	}
	if teacherId != ""{
		query = query.Where("teacherid = ?", teacherId)
	}

	err := query.Find(&schedules).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(schedules)
}

// func ScheduleFilter(typeFilter, queryString string) *gorm.DB{
// 	if typeFilter != ""{
// 		return config.DB.Model(&Schedule{}).Where(queryString, typeFilter)
// 	}
// 	return
// }

func UpdateScheduleHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var schedule Schedule

	err := config.DB.First(&schedule, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Расписание не найдено", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	var updatedData Schedule
	err = json.NewDecoder(r.Body).Decode(&updatedData)
	if err != nil {
		http.Error(w, "Неверное тело запроса", http.StatusBadRequest)
		return
	}

	err = config.DB.Model(&schedule).Updates(updatedData).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(schedule)
}

func AddBulkSchedulesHandler(w http.ResponseWriter, r *http.Request) {
	var newSchedules []NewSchedule

	err := json.NewDecoder(r.Body).Decode(&newSchedules)
	if err != nil {
		http.Error(w, "Неверное тело запроса", http.StatusBadRequest)
		return
	}

	var schedules []Schedule

	for _, schedule := range newSchedules {
		newSchedule := Schedule{
			GroupID:    schedule.GroupID,
			SubjectID:  schedule.SubjectID,
			IsOccupied: false,
		}
		if schedule.IsOccupied != nil {
			newSchedule.IsOccupied = *schedule.IsOccupied
		}
		schedules = append(schedules, newSchedule)
	}

	err = config.DB.Create(&schedules).Error
	if err != nil {
		http.Error(w, "Ошибка при добавлении расписания", http.StatusInternalServerError)
		return
	}
	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(schedules)
}

func DeleteScheduleHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := config.DB.Delete(&Schedule{}, id).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteBulkSchedulesHandler(w http.ResponseWriter, r *http.Request){
	var schedulesID [] uint
	err := json.NewDecoder(r.Body).Decode(&schedulesID)
	if err != nil{
		http.Error(w, "Неверное тело запроса", http.StatusBadRequest)
		return
	}
	if len(schedulesID) == 0 {
		http.Error(w, "Пустое тело запроса", http.StatusBadRequest)
		return
	}
	result := config.DB.Where("scheduleid IN ?", schedulesID).Delete(&Schedule{})
	if result.Error != nil{
		http.Error(w, "Ошибка удаления указанных расписаний", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message": "Расписания удалены успешно",
		"count": result.RowsAffected,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
