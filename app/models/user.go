package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID      string `gorm:"primary_key;type:varchar(100);not null" json:"id"`
	Name    string `gorm:"type:varchar(50);not null" json:"name"`
	Surname string `gorm:"type:varchar(50);not null" json:"surname"`
}

type UserResponse struct {
	ID        string     `json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
	Name      string     `json:"name"`
	Surname   string     `json:"surname"`
}
