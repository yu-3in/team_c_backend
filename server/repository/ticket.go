package repository

import (
	"server/model"
	"time"
)

func (r *Repository) GetTickets() ([]*model.Ticket, error) {
	// var tickets []*model.Ticket
	// result := r.db.Find(&tickets)
	// if result.Error != nil {
	// 	return nil, result.Error
	// }

	tickets := []*model.Ticket{
		{
			ID:          1,
			Title:       "Make CSS",
			Status:      "Do",
			DueDate:     time.Now(),
			StartAt:     time.Now(),
			EndAt:       time.Now(),
			Description: "test",
			User: model.User{
				ID:   1,
				Name: "John",
			},
			Genre: model.Genre{
				ID:        1,
				Title:     "BackEnd#1",
				Color:     "Red",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:          2,
			Title:       "Make API",
			Status:      "Done",
			DueDate:     time.Now(),
			StartAt:     time.Now(),
			EndAt:       time.Now(),
			Description: "test",
			User: model.User{
				ID:   1,
				Name: "John",
			},
			Genre: model.Genre{
				ID:        1,
				Title:     "BackEnd#1",
				Color:     "Red",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	return tickets, nil
}

func (r *Repository) GetTicket(id int) (*model.Ticket, error) {
	// var ticket model.Ticket
	// result := r.db.First(&ticket, id)
	// if result.Error != nil {
	// 	return nil, result.Error
	// }

	ticket := model.Ticket{
		ID:          1,
		Title:       "Make CSS",
		DueDate:     time.Now(),
		StartAt:     time.Now(),
		EndAt:       time.Now(),
		Description: "test",
		User: model.User{
			ID:   1,
			Name: "John",
		},
		Genre: model.Genre{
			ID:        1,
			Title:     "BackEnd#1",
			Color:     "Red",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return &ticket, nil
}
