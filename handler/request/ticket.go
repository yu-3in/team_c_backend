package request

import (
	"time"
)

type ReqGetTicket struct {
	GenreID int `query:"genreId"`
	UserID  int `query:"userId"`

	Sort string `query:"sort"`
	Reco []int 
}

type ReqCreateTicket struct {
	Title       string    `json:"title"`
	Status      string    `json:"status"`
	DueDate     time.Time `json:"dueDate"`
	StartAt     time.Time `json:"startAt"`
	EndAt       time.Time `json:"endAt"`
	Description string    `json:"description"`
	UserID      int       `json:"userId"`
	GenreID     int       `json:"genreId"`
}

type ReqUpdateTicket struct {
	ID               int       `param:"id"`
	Title            string    `json:"title"`
	Status           string    `json:"status"`
	DueDate          time.Time `json:"dueDate"`
	StartAt          time.Time `json:"startAt"`
	EndAt            time.Time `json:"endAt"`
	Description      string    `json:"description"`
	RaisedHandUserId int       `json:"raisedHandUserId"`
	UserID           int       `json:"userId"`
	GenreID          int       `json:"genreId"`
}
