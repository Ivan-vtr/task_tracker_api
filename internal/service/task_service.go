package service

import (
	"context"
	"errors"
	"task_tracker_api/internal/model"
	"task_tracker_api/internal/repository"
)

type TaskService interface {
	Create(ctx context.Context, task *model.Task) error
	Get(ctx context.Context, id int64) (*model.Task, error)
	GetAll(ctx context.Context) ([]*model.Task, error)
}

type taskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) TaskService {
	return &taskService{repo: repo}
}

func (s *taskService) Create(ctx context.Context, task *model.Task) error {
	if task.Title == "" {
		return errors.New("title is required")
	}
	task.Status = "new"
	task.UserID = 1

	return s.repo.Create(ctx, task)
}

func (s *taskService) Get(
	ctx context.Context,
	id int64,
) (*model.Task, error) {
	return s.repo.GetById(ctx, id)
}

func (s *taskService) GetAll(ctx context.Context) ([]*model.Task, error) {
	return s.repo.GetAll(ctx)
}
