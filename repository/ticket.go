package repository

import (
	"server/handler/request"
	"server/model"
)

func (r *Repository) GetTickets(req request.ReqGetTicket) ([]*model.Ticket, error) {
	var tickets []*model.Ticket

	query := r.db.Preload("User").Preload("Genre")

	if req.GenreID != 0 {
		query = query.Where("genre_id = ?", req.GenreID)
	}
	
	if req.UserID != 0 {
		query = query.Where("user_id = ?", req.UserID)
	}

	switch req.Sort {
	case "latest_creation":
		query = query.Order("created_at desc")
	case "oldest_creation":
		query = query.Order("created_at asc")
	case "latest_update":
		query = query.Order("updated_at desc")
	case "oldest_update":
		query = query.Order("updated_at asc")
	}

	result := query.Find(&tickets)
	if result.Error != nil {
		return nil, result.Error
	}

	if req.Sort == "recommended" {
		filteredTickets := filterTicketsByGenre(tickets, req.Reco)
		return filteredTickets, nil
	}
	return tickets, nil
}

func filterTicketsByGenre(tickets []*model.Ticket, genreIDs []int) []*model.Ticket {
	var filteredTickets []*model.Ticket
	for _, ticket := range tickets {
		for _, genreID := range genreIDs {
			if ticket.GenreID == genreID {
				filteredTickets = append(filteredTickets, ticket)
				break
			}
		}
	}
	return filteredTickets
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
	sql := `UPDATE tickets SET title=?, status=?, due_date=?, start_at=?, end_at=?, description=?, user_id=?, genre_id=?, created_at=?, updated_at=? WHERE id=?`
	resuelt := r.db.Exec(sql, ticket.Title, ticket.Status, ticket.DueDate, ticket.StartAt, ticket.EndAt, ticket.Description, ticket.UserID, ticket.GenreID, ticket.CreatedAt, ticket.UpdatedAt, ticket.ID)

	if resuelt.Error != nil {
		return nil, resuelt.Error
	}
	return ticket, nil
}

func (r *Repository) DeleteTicket(id int) error {
	result := r.db.Delete(&model.Ticket{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
