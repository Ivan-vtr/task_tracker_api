package model

import "time"

type User struct {
	ID           uint      `gorm:"primaryKey"`
	Email        string    `gorm:"size:100;unique;not null"`
	PasswordHash string    `gorm:"size:255;not null"`
	Role         string    `gorm:"size:50;not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
}
