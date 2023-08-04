package handler

import (
	"server/handler/request"
	"server/model"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetTickets(c echo.Context) error {
	var req request.ReqGetTicket
	if err := c.Bind(&req); err != nil {
		return err
	}
	tickets, err := h.repo.GetTickets(req)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, tickets)
}

func (h *Handler) GetTicket(c echo.Context) error {
	var req struct {
		ID int `param:"id"`
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

func (h *Handler) CreateTicket(c echo.Context) error {
	var req request.ReqCreateTicket
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

func (h *Handler) UpdateTicket(c echo.Context) error {
	var req request.ReqUpdateTicket
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
	if req.RaisedHandUserId != 0 {
		ticket.RaisedHandUserId = req.RaisedHandUserId
	} else {
		ticket.RaisedHandUserId = 0
	}
	if req.UserID != 0 {
		ticket.UserID = req.UserID
	} else {
		ticket.UserID = 0
	}
	ticket.GenreID = req.UserID

	ticket, err = h.repo.UpdateTicket(ticket)
	if err != nil {
		return c.JSON(500, err)
	}

	return c.JSON(200, ticket)
}

func (h *Handler) DeleteTicket(c echo.Context) error {
	var req struct {
		ID int `param:"id"`
	}
	if err := c.Bind(&req); err != nil {
		return err
	}
	err := h.repo.DeleteTicket(req.ID)
	if err != nil {
		return err
	}
	return c.NoContent(200)
}
