package service

import (
	"Checkmarx/model"
	"Checkmarx/repository"
)

type TaskService struct {
	repo *repository.TaskManager
}

func NewTaskService(repo *repository.TaskManager) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) AddTask(task model.Task) (model.Task, error) {
	return s.repo.AddTask(task)
}

func (s *TaskService) GetTask(id int) (model.Task, error) {
	return s.repo.GetTask(id)
}

func (s *TaskService) GetAllTasks() map[int]model.Task {
	return s.repo.GetAllTasks()
}

func (s *TaskService) UpdateTask(id int, updatedTask model.Task) (model.Task, error) {
	return s.repo.Update(id, updatedTask)
}

func (s *TaskService) DeleteTask(id int) error {
	return s.repo.Delete(id)
}
