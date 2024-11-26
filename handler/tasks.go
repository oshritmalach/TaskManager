package handler

import (
	"Checkmarx/helpers"
	"Checkmarx/model"
	"Checkmarx/service"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type TaskHandler struct {
	service *service.TaskService
}

func NewTaskHandler(service *service.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

func (h *TaskHandler) AddTask(w http.ResponseWriter, r *http.Request) {
	var err error
	var newTask model.Task
	err = json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Invalid request body")
		return
	}

	if valid, errValidation := helpers.ValidateTaskFields(newTask); !valid {
		respondWithError(w, http.StatusBadRequest, errValidation.Error())
		return
	}
	if _, err = h.service.AddTask(newTask); err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (h *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	id, err := getIdParam(r)
	if err != nil || id <= 0 {
		respondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	task, err := h.service.GetTask(id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(task); err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to encode task:%d,%v", id, err))
		return
	}
}

func (h *TaskHandler) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks := h.service.GetAllTasks()
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to encode all tasks: %v", err))
		return
	}
}

func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	id, err := getIdParam(r)
	if err != nil || id <= 0 {
		respondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	var updatedTask model.Task
	if err = json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Invalid request body")
		return
	}

	var valid bool
	if valid, err = helpers.ValidateTaskUpdate(updatedTask); !valid {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	task, err := h.service.UpdateTask(id, updatedTask)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(task); err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to encode task:%d,%v", id, err))
		return
	}
}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id, err := getIdParam(r)
	if err != nil || id <= 0 {
		respondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	if err = h.service.DeleteTask(id); err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func getIdParam(r *http.Request) (int, error) {
	var id int
	var err error
	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr != "" {
		id, err = strconv.Atoi(idStr)
		return id, err
	}
	return id, nil
}

func respondWithError(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	errorResponse := map[string]interface{}{
		"message": message,
		"code":    statusCode,
	}

	if err := json.NewEncoder(w).Encode(errorResponse); err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode error response: %v", err), http.StatusInternalServerError)
		return
	}
}
