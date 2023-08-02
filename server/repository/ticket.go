package repository

import (
	"server/model"
)

func (r *Repository) GetTickets() ([]*model.Ticket, error) {
	var tickets []*model.Ticket
	result := r.db.Preload("User").Preload("Genre").First(&tickets)
	if result.Error != nil {
		return nil, result.Error
	}
	return tickets, nil
}

func (r *Repository) GetTicket(id int) (*model.Ticket, error) {
	var ticket model.Ticket
	result := r.db.Preload("User").Preload("Genre").First(&ticket, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &ticket, nil
}

func (r *Repository) CreateTicket(ticket *model.Ticket) (*model.Ticket, error) {
	result := r.db.Create(ticket)
	if result.Error != nil {
		return nil, result.Error
	}
	return ticket, nil
}

func (r *Repository) UpdateTicket(ticket *model.Ticket) (*model.Ticket, error) {
	result := r.db.Save(ticket)
	if result.Error != nil {
		return nil, result.Error
	}
	return ticket, nil
}
