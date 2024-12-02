package httputil

import (
	"net/http"
	"teacher_schedule/internal/groups"
	"teacher_schedule/internal/schedule"
	"teacher_schedule/internal/subjects"

	
	"teacher_schedule/internal/users"

	"github.com/go-chi/chi/v5"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("АПИ Расписание Преподавателя"))
	})

	r.Route("/users", func(r chi.Router) {
		r.Get("/", users.GetUserHandler)
		r.Post("/register", users.CreateUserHandler)
		r.Get("/teachers", users.GetTeachersHandler)
	})
	r.Route("/schedule", func(r chi.Router) {
		r.Get("/", schedule.GetAllSchedulesHandler)
		r.Get("/{id}", schedule.GetScheduleByIdHandler)
		r.Get("/teacher/{teacherId}", schedule.GetScheduleByTeacherHandler)
		r.Get("/filter", schedule.GetScheduleFilteredHandler)
		r.Post("/", schedule.CreateScheduleHandler)
		r.Put("/{id}", schedule.UpdateScheduleHandler)
		r.Post("/bulk-add", schedule.BulkAddSchedulesHandler)
		r.Delete("/{id}", schedule.DeleteScheduleHandler)
	})
	r.Route("/subjects", func(r chi.Router) {
		r.Get("/", subjects.GetAllSubjectsHandler)
		r.Post("/", subjects.CreateSunbjectHandler)
	})
	r.Route("/group", func(r chi.Router){
		r.Get("/", groups.GetGroupHandler)
		r.Post("/", groups.CreateGroupHandler)
	})

	return r
}
