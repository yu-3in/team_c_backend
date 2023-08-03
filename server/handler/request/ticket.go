package request

import (
	"time"
)

type ReqGetTicket struct {
	Genre string `query:"genre"`
}

type ReqCreateTicket struct {
	Title       string    `json:"title"`
	Status      string    `json:"status"`
	DueDate     time.Time `json:"dateDate"`
	StartAt     time.Time `json:"startAt"`
	EndAt       time.Time `json:"endAt"`
	Description string    `json:"description"`
	UserID      int       `json:"userId"`
	GenreID     int       `json:"genreId"`
}

type ReqUpdateTicket struct {
	ID          int       `param:"id"`
	Title       string    `json:"title"`
	Status      string    `json:"status"`
	DueDate     time.Time `json:"dateDate"`
	StartAt     time.Time `json:"startAt"`
	EndAt       time.Time `json:"endAt"`
	Description string    `json:"description"`
	UserID      int       `json:"userId"`
	GenreID     int       `json:"genreId"`
}
