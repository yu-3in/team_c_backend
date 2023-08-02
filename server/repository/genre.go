package repository

import (
	"server/model"
)

func (r *Repository) GetGenres() ([]*model.Genre, error) {
	var genres []*model.Genre
	result := r.db.Find(&genres)
	if result.Error != nil {
		return nil, result.Error
	}
	return genres, nil
}

func (r *Repository) GetGenre(id int) (*model.Genre, error) {
	var genre model.Genre
	result := r.db.First(&genre)
	if result.Error != nil {
		return nil, result.Error
	}
	return &genre, nil
}

func (r *Repository) CreateGenre(genre *model.Genre) (*model.Genre, error) {
	result := r.db.Create(genre)
	if result.Error != nil {
		return nil, result.Error
	}
	return genre, nil
}

func (r *Repository) UpdateGenre(genre *model.Genre) (*model.Genre, error) {
	result := r.db.Save(genre)
	if result.Error != nil {
		return nil, result.Error
	}
	return genre, nil
}
