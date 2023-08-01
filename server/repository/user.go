package repository

import (
	"server/model"
	"time"
)

func (r *Repository) GetUsers() ([]*model.User, error) {
	// var users []*model.User
	// result := r.db.Find(&users)
	// if result.Error != nil {
	// 	return nil, result.Error
	// }

	users := []*model.User{
		{
			ID:       1,
			Name:     "John",
			Email:    "exsample@exsample.com",
			Password: "password",
			Genres: []model.Genre{
				{
					ID:        1,
					Title:     "BackEnd",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
			},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:       2,
			Name:     "Mike",
			Email:    "exsample2@exsample.com",
			Password: "password",
			Genres: []model.Genre{
				{
					ID:        1,
					Title:     "FrontEnd",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
			},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	return users, nil
}

func (r *Repository) GetUser(id int) (*model.User, error) {
	// var user model.User
	// result := r.db.First(&user, id)
	// if result.Error != nil {
	// 	return nil, result.Error
	// }

	user := model.User{
		ID:       1,
		Name:     "John",
		Email:    "exsample@exsample.com",
		Password: "password",
		Genres: []model.Genre{
			{
				ID:        1,
				Title:     "BackEnd",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return &user, nil
}
