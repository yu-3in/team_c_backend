package repository

import (
	"math/rand"
	"server/handler/request"
	"server/model"
	"time"
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
		filteredTickets := filterAndCombineTicketsByGenre(tickets, req.Reco)
		return filteredTickets, nil
	}
	return tickets, nil
}

func filterAndCombineTicketsByGenre(tickets []*model.Ticket, genreIDs []int) []*model.Ticket {
	var filteredTickets []*model.Ticket
	var unfilteredTickets []*model.Ticket

	for _, ticket := range tickets {
		matched := false
		for _, genreID := range genreIDs {
			if ticket.GenreID == genreID {
				filteredTickets = append(filteredTickets, ticket)
				matched = true
				break
			}
		}
		if !matched {
			unfilteredTickets = append(unfilteredTickets, ticket)
		}
	}

	shuffleTickets(filteredTickets)
	resultTickets := append(filteredTickets, unfilteredTickets...)
	return resultTickets
}

func shuffleTickets(tickets []*model.Ticket) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := len(tickets) - 1; i > 0; i-- {
		j := r.Intn(i + 1)
		tickets[i], tickets[j] = tickets[j], tickets[i]
	}
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
	var userID, genreID interface{}
	if ticket.UserID == 0 {
		userID = nil
	} else {
		userID = ticket.UserID
	}
	
	if ticket.GenreID == 0 {
		genreID = nil
	} else {
		genreID = ticket.GenreID
	}
	sql := `UPDATE tickets SET title=?, status=?, due_date=?, start_at=?, end_at=?, description=?, user_id=?, genre_id=?, created_at=?, updated_at=? WHERE id=?`
	result := r.db.Exec(sql, ticket.Title, ticket.Status, ticket.DueDate, ticket.StartAt, ticket.EndAt, ticket.Description, userID, genreID, ticket.CreatedAt, ticket.UpdatedAt, ticket.ID)

	if result.Error != nil {
		return nil, result.Error
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
