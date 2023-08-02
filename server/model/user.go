package model

import (
	"time"
)

type User struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"unique"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Genres    []Genre   `json:"genres" gorm:"many2many:user_genres;"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
