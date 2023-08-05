package request

import (
	"server/model"
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

type ResTicket struct {
	ID             int       `json:"id"`
	Title          string    `json:"title"`
	Status         string    `json:"status"`
	DueDate        time.Time `json:"dueDate"`
	StartAt        time.Time `json:"startAt"`
	EndAt          time.Time `json:"endAt"`
	Description    string    `json:"description"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	User           *ResUser
	Genre          *ResGenre
	RaisedHandUser *ResUser
}

func TicketModelToResTicket(m *model.Ticket) *ResTicket {
	return &ResTicket{
		ID:             m.ID,
		Title:          m.Title,
		Status:         m.Status,
		StartAt:        m.StartAt,
		EndAt:          m.EndAt,
		Description:    m.Description,
		User:           UserModelToResUser(&m.User),
		Genre:          GenreModelToResGenre(&m.Genre),
		RaisedHandUser: UserModelToResUser(&m.RaisedHandUser),
		CreatedAt:      m.CreatedAt,
		UpdatedAt:      m.UpdatedAt,
	}
}

func TicketModelToResTickets(m []*model.Ticket) []*ResTicket {
	responses := make([]*ResTicket, len(m))
	for i, ticket := range m {
		responses[i] = TicketModelToResTicket(ticket)
	}
	return responses
}
