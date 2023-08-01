package repository

import (
	"server/model"
	"time"
)

func (r *Repository) GetGenres() ([]*model.Genre, error) {
	// var genres []*model.Genre
	// result := r.db.Find(&genres)
	// if result.Error != nil {
	// 	return nil, result.Error
	// }
	genres := []*model.Genre{
		{
			ID:        1,
			Title:     "BackEnd#1",
			Color:     "Red",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        1,
			Title:     "FrontEnd#1",
			Color:     "Blue",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	return genres, nil
}

func (r *Repository) GetGenre(id int) (*model.Genre, error) {
	// var genre model.Genre
	// result := r.db.First(&genre)
	// if result.Error != nil {
	// 	return nil, result.Error
	// }

	genre := model.Genre{
		ID:        1,
		Title:     "BackEnd#1",
		Color:     "Red",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return &genre, nil
}
