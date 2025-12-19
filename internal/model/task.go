package model

import "time"

type Task struct {
	ID          uint      `gorm:"primaryKey"`
	UserID      uint      `gorm:"not null"` // внешний ключ на users.id
	Title       string    `gorm:"size:200;not null"`
	Description string    `gorm:"type:text"`
	Status      string    `gorm:"size:50;default:'pending'"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
