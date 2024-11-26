package handler

import (
	"Checkmarx/model"
	"Checkmarx/repository"
	"Checkmarx/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func startTestServer() *mux.Router {

	repo := repository.NewTaskManager()
	service := service.NewTaskService(repo)
	handler := TaskHandler{service: service}
	r := mux.NewRouter()

	r.HandleFunc("/task", handler.AddTask).Methods("POST")
	r.HandleFunc("/tasks", handler.GetAllTasks).Methods("GET")
	r.HandleFunc("/task/{id}", handler.GetTask).Methods("GET")
	r.HandleFunc("/task/{id}", handler.UpdateTask).Methods("POST")
	r.HandleFunc("/task/{id}", handler.DeleteTask).Methods("DELETE")
	return r
}

func TestAddTaskHandler(t *testing.T) {
	server := httptest.NewServer(startTestServer())
	defer server.Close()

	requestBody := `{
        "title": "New Task",
        "description": "Description of the task",
        "status": "pending"
    }`

	req, err := http.NewRequest("POST", server.URL+"/task", strings.NewReader(requestBody))
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("expected status code %d, got %d", http.StatusCreated, resp.StatusCode)
	}
}
func TestGetAllTasksHandler(t *testing.T) {
	server := httptest.NewServer(startTestServer())
	defer server.Close()

	req, err := http.NewRequest("GET", server.URL+"/tasks", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}
}
func TestGetTaskIfNotExistsHandler(t *testing.T) {
	server := httptest.NewServer(startTestServer())
	defer server.Close()

	req, err := http.NewRequest("GET", server.URL+"/task/1", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("expected status code %d, got %d", http.StatusNotFound, resp.StatusCode)
	}
}
func TestAddAndGetTaskHandler(t *testing.T) {
	server := httptest.NewServer(startTestServer())
	defer server.Close()

	requestBody := `{"title": "New Task", "description": "Task description", "status": "pending"}`
	req, err := http.NewRequest("POST", server.URL+"/task", strings.NewReader(requestBody))
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("expected status code %d, got %d", http.StatusCreated, resp.StatusCode)
	}

	req, err = http.NewRequest("GET", server.URL+"/task/1", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	// 4. קריאת התשובה והשוואת הערכים
	var task model.Task
	if err := json.NewDecoder(resp.Body).Decode(&task); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	// בודק שהערכים שהתקבלו תואמים למה שנשלח
	if task.Title != "New Task" || task.Description != "Task description" || task.Status != "pending" {
		t.Errorf("expected task details to be 'New Task', 'Task description', 'pending', but got '%s', '%s', '%s'", task.Title, task.Description, task.Status)
	}
}
func TestDeleteIfNotExistsHandler(t *testing.T) {
	server := httptest.NewServer(startTestServer())
	defer server.Close()

	req, err := http.NewRequest("DELETE", server.URL+"/task/1", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("expected status code %d, got %d", http.StatusNotFound, resp.StatusCode)
	}
}
func TestAddAndDeleteTaskHandler(t *testing.T) {
	server := httptest.NewServer(startTestServer())
	defer server.Close()

	requestBody := `{"title": "New Task", "description": "Task description", "status": "pending"}`
	req, err := http.NewRequest("POST", server.URL+"/task", strings.NewReader(requestBody))
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("expected status code %d, got %d", http.StatusCreated, resp.StatusCode)
	}

	req, err = http.NewRequest("DELETE", server.URL+"/task/1", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("expected status code %d, got %d", http.StatusNoContent, resp.StatusCode)
	}
}
func TestUpdateIfNotExistsHandler(t *testing.T) {
	server := httptest.NewServer(startTestServer())
	defer server.Close()

	requestBody := `{"title": "New Task", "description": "Task description", "status": "pending"}`
	req, err := http.NewRequest("POST", server.URL+"/task/1", strings.NewReader(requestBody))
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("expected status code %d, got %d", http.StatusNotFound, resp.StatusCode)
	}
}
func TestAddAndUpdateTaskHandler(t *testing.T) {
	server := httptest.NewServer(startTestServer())
	defer server.Close()

	requestBody := `{"title": "New Task", "description": "Task description", "status": "pending"}`
	req, err := http.NewRequest("POST", server.URL+"/task", strings.NewReader(requestBody))
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("expected status code %d, got %d", http.StatusCreated, resp.StatusCode)
	}

	UpdatedRequestBody := `{"title": "New Task", "description": "Task description", "status": "pending"}`
	req, err = http.NewRequest("POST", server.URL+"/task/1", strings.NewReader(UpdatedRequestBody))
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}
}
