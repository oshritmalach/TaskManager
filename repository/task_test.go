package repository

import (
	"Checkmarx/helpers"
	"Checkmarx/model"
	"testing"
	"time"
)

func TestTaskRepository(t *testing.T) {
	repo := NewTaskManager()

	newTaskPositive := model.Task{
		Title:       "Test Task",
		Description: "This is a test task.",
		Status:      "Pending",
		CreatedAt:   time.Now(),
	}

	_, err := repo.AddTask(newTaskPositive)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if repo.counter != 1 {
		t.Errorf("expected ID 1, got %d", repo.counter)
	}

	task, err := repo.GetTask(repo.counter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if task.Title != "Test Task" {
		t.Errorf("expected title 'Test Task', got '%s'", task.Title)
	}

	updatedTaskData := model.Task{
		Title:       "Updated Task",
		Description: "This is a test task.",
		Status:      "Completed",
	}
	updatedTask, err := repo.Update(repo.counter, updatedTaskData)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if updatedTask.Title != "Updated Task" || updatedTask.Status != "Completed" || updatedTask.Description != "This is a test task." {
		t.Errorf("update failed: %+v", updatedTask)
	}

	allTasks := repo.GetAllTasks()
	if len(allTasks) != 1 {
		t.Errorf("expected 1 task, got %d", len(allTasks))
	}

	err = repo.Delete(repo.counter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	_, err = repo.GetTask(repo.counter)
	if err == nil {
		t.Errorf("expected error, got none")
	}

}

func TestValidateAddTask(t *testing.T) {
	tests := []struct {
		name     string
		task     model.Task
		expected bool
		errMsg   string
	}{
		{
			name: "Valid task",
			task: model.Task{
				Title:       "Test Task",
				Description: "Test Description",
				Status:      "open",
				CreatedAt:   time.Now(),
			},
			expected: true,
			errMsg:   "",
		},
		{
			name: "Empty title",
			task: model.Task{
				Title:       "",
				Description: "Test Description",
				Status:      "open",
				CreatedAt:   time.Now(),
			},
			expected: false,
			errMsg:   "title is required",
		},
		{
			name: "Empty description",
			task: model.Task{
				Title:       "Test Task",
				Description: "",
				Status:      "open",
				CreatedAt:   time.Now(),
			},
			expected: false,
			errMsg:   "description is required",
		},
		{
			name: "Empty status",
			task: model.Task{
				Title:       "Test Task",
				Description: "Test Description",
				Status:      "",
				CreatedAt:   time.Now(),
			},
			expected: false,
			errMsg:   "status is required",
		},
		{
			name: "Empty title, description and status",
			task: model.Task{
				Title:       "",
				Description: "",
				Status:      "",
				CreatedAt:   time.Now(),
			},
			expected: false,
			errMsg:   "title is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid, err := helpers.ValidateTaskFields(tt.task)
			if valid != tt.expected {
				t.Errorf("expected valid: %v, got %v", tt.expected, valid)
			}
			if err != nil && err.Error() != tt.errMsg {
				t.Errorf("expected error message: %v, got: %v", tt.errMsg, err.Error())
			}
		})
	}
}

func TestValidateUpdateTask(t *testing.T) {
	tests := []struct {
		name     string
		task     model.Task
		expected bool
		errMsg   string
	}{
		{
			name: "Valid task",
			task: model.Task{
				Title:       "Test Task",
				Description: "Test Description",
				Status:      "open",
				CreatedAt:   time.Now(),
			},
			expected: true,
			errMsg:   "",
		},
		{
			name: "Empty title, description and status",
			task: model.Task{
				Title:       "",
				Description: "",
				Status:      "",
				CreatedAt:   time.Now(),
			},
			expected: false,
			errMsg:   "at least one field (Title, Description, Status) must be updated",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid, err := helpers.ValidateTaskUpdate(tt.task)
			if valid != tt.expected {
				t.Errorf("expected valid: %v, got %v", tt.expected, valid)
			}
			if err != nil && err.Error() != tt.errMsg {
				t.Errorf("expected error message: %v, got: %v", tt.errMsg, err.Error())
			}
		})
	}
}
