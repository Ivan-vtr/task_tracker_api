package repository

import (
	"context"
	"task_tracker_api/internal/model"

	"github.com/jmoiron/sqlx"
)

type TaskRepository interface {
	Create(ctx context.Context, task *model.Task) error
	GetById(ctx context.Context, id int64) (*model.Task, error)
}

type taskRepository struct {
	db *sqlx.DB
}

func NewTaskRepository(db *sqlx.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) Create(
	ctx context.Context,
	task *model.Task,
) error {
	//Todo add real active user_id
	query := `
		INSERT INTO tasks (title, description, status, user_id)
		VALUES ($1, $2, $3, 1)
		RETURNING id, created_at
	`

	return r.db.QueryRowxContext(
		ctx,
		query,
		task.Title,
		task.Description,
		task.Status,
	).Scan(&task.ID, &task.CreatedAt)
}

func (r *taskRepository) GetById(
	ctx context.Context,
	id int64,
) (*model.Task, error) {
	var task model.Task

	query := `
		SELECT id, created_at, title, description, status
		FROM tasks
		WHERE id = $1
		`
	if err := r.db.GetContext(ctx, &task, query, id); err != nil {
		return nil, err
	}

	return &task, nil
}
