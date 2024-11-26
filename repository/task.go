package repository

import (
	"Checkmarx/helpers"
	"Checkmarx/model"
	"errors"
	"fmt"
	"sync"
	"time"
)

type TaskManager struct {
	mu      sync.RWMutex
	tasks   map[int]model.Task
	counter int
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		tasks:   make(map[int]model.Task),
		counter: 0,
	}
}

func (tm *TaskManager) AddTask(newTask model.Task) (model.Task, error) {
	if valid, err := helpers.ValidateTaskFields(newTask); !valid {
		return model.Task{}, err
	}

	task := model.Task{
		Title:       newTask.Title,
		Description: newTask.Description,
		Status:      newTask.Status,
		CreatedAt:   time.Now(),
	}

	tm.mu.Lock()
	tm.counter++
	tm.tasks[tm.counter] = task
	tm.mu.Unlock()
	return task, nil
}

func (tm *TaskManager) Update(id int, updatedTask model.Task) (model.Task, error) {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	task, exists := tm.tasks[id]
	if !exists {
		return model.Task{}, errors.New(fmt.Sprintf("task number %d does not exist", id))
	}

	if valid, err := helpers.ValidateTaskUpdate(updatedTask); !valid {
		return model.Task{}, err
	}

	if updatedTask.Title != "" {
		task.Title = updatedTask.Title
	}
	if updatedTask.Description != "" {
		task.Description = updatedTask.Description
	}
	if updatedTask.Status != "" {
		task.Status = updatedTask.Status
	}

	tm.tasks[id] = task
	return tm.tasks[id], nil
}

func (tm *TaskManager) GetAllTasks() map[int]model.Task {
	tm.mu.RLock()
	defer tm.mu.RUnlock()
	return tm.tasks
}

func (tm *TaskManager) Delete(id int) error {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	if _, exists := tm.tasks[id]; !exists {
		return fmt.Errorf("task number %d does not exist", id)

	}
	delete(tm.tasks, id)
	return nil
}

func (tm *TaskManager) GetTask(id int) (model.Task, error) {
	tm.mu.RLock()
	defer tm.mu.RUnlock()
	task, exists := tm.tasks[id]
	if !exists {
		return model.Task{}, fmt.Errorf("task number %d does not exist", id)
	}
	return task, nil
}
