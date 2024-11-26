package router

import (
	"Checkmarx/handler"
	"Checkmarx/repository"
	"Checkmarx/service"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	repo := repository.NewTaskManager()
	service := service.NewTaskService(repo)
	handler := handler.NewTaskHandler(service)

	r := mux.NewRouter()
	r.HandleFunc("/tasks", handler.GetAllTasks).Methods("GET")
	r.HandleFunc("/task", handler.AddTask).Methods("POST")
	r.HandleFunc("/task/{id}", handler.UpdateTask).Methods("POST")
	r.HandleFunc("/task/{id}", handler.DeleteTask).Methods("DELETE")
	r.HandleFunc("/task/{id}", handler.GetTask).Methods("GET")
	return r
}
