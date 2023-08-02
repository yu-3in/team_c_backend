package handler

import (
	"server/model"
	"time"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetTickets(c echo.Context) error {
	tickets, err := h.repo.GetTickets()
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, tickets)
}

func (h *Handler) GetTicket(c echo.Context) error {
	var req struct {
		ID int `json:"id"`
	}
	if err := c.Bind(&req); err != nil {
		return err
	}
	ticket, err := h.repo.GetTicket(req.ID)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, ticket)
}

type reqCreateTicket struct {
	Title       string    `json:"title"`
	Status      string    `json:"status"`
	DueDate     time.Time `json:"dateDate"`
	StartAt     time.Time `json:"startAt"`
	EndAt       time.Time `json:"endAt"`
	Description string    `json:"description"`
	UserID      int       `json:"userId"`
	GenreID     int       `json:"genreId"`
}

func (h *Handler) CreateTicket(c echo.Context) error {
	var req reqCreateTicket
	if err := c.Bind(&req); err != nil {
		return err
	}

	ticket, err := h.repo.CreateTicket(&model.Ticket{
		Title:       req.Title,
		Status:      req.Status,
		DueDate:     req.DueDate,
		StartAt:     req.StartAt,
		EndAt:       req.EndAt,
		Description: req.Description,
		UserID:      req.UserID,
		GenreID:     req.GenreID,
	})
	if err != nil {
		return c.JSON(500, err)
	}

	return c.JSON(200, ticket)
}

type reqUpdateTicket struct {
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

func (h *Handler) UpdateTicket(c echo.Context) error {
	var req reqUpdateTicket
	if err := c.Bind(&req); err != nil {
		return err
	}

	ticket, err := h.repo.GetTicket(req.ID)
	if err != nil {
		return c.JSON(500, err)
	}

	ticket.Title = req.Title
	ticket.Status = req.Status
	ticket.DueDate = req.DueDate
	ticket.StartAt = req.StartAt
	ticket.EndAt = req.EndAt
	ticket.Description = req.Description
	ticket.UserID = req.UserID
	ticket.GenreID = req.GenreID

	ticket, err = h.repo.UpdateTicket(ticket)
	if err != nil {
		return c.JSON(500, err)
	}

	return c.JSON(200, ticket)
}
