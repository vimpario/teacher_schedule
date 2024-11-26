package httputil

import (
	"net/http"
	"teacher_schedule/internal/users"

	"github.com/go-chi/chi/v5"
)

func NewRouter() *chi.Mux{
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("АПИ Расписание Преподавателя"))
	})

	r.Route("/users", func(r chi.Router){
		r.Get("/", users.GetUserHandler)
		r.Post("/", users.CreateUserHandler)
	})

	return r
}