package repository

import (
	"fmt"
	"server/handler/request"
	"server/model"
)

// Sort
// 'latest_creation' // 作成日の新しい順
// 'oldest_creation' // 古い順
// 'latest_update' // 更新の新しい順
// 'oldest_update' // 更新の古い順

//filter
// genreID
// userID

func (r *Repository) GetTickets(req request.ReqGetTicket) ([]*model.Ticket, error) {
	var tickets []*model.Ticket

	query := r.db.Preload("User").Preload("Genre")

	if req.GenreID != 0 {
		fmt.Println("genreID", req.GenreID)
		query = query.Where("genre_id = ?", req.GenreID)
	}
	
	if req.UserID != 0 {
		query = query.Where("user_id = ?", req.UserID)
	}

	// if req.Sort != "" {
	// 	query = query.Order(req.Sort) // ソート順を指定
	// }

	result := query.Find(&tickets)

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
