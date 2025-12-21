package model

import "time"

type Task struct {
	ID          uint      `db:"id"`
	UserID      uint      `db:"user_id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	Status      string    `db:"status"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
