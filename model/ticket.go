package model

import (
	"time"
)

type Ticket struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title"`
	Status      string    `json:"status"`
	DueDate     time.Time `json:"dueDate"`
	StartAt     time.Time `json:"startAt"`
	EndAt       time.Time `json:"endAt"`
	Description string    `json:"description"`
	UserID      int       `json:"userId"`
	GenreID     int       `json:"genreId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	User        User      `gorm:"foreignKey:UserID"`
	Genre       Genre     `gorm:"foreignKey:GenreID"`
}
