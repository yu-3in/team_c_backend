package repository

import (
	"server/model"
)

func (r *Repository) GetUsers() ([]*model.User, error) {
	var users []*model.User
	result := r.db.Preload("Genres").Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (r *Repository) GetUser(id int) (*model.User, error) {
	var user model.User
	result := r.db.Preload("Genres").First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *Repository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	result := r.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *Repository) CreateUser(user *model.User) (*model.User, error) {
	result := r.db.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (r *Repository) UpdateUser(user *model.User) (*model.User, error) {
	result := r.db.Save(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (r *Repository) CreateUserGenre(user *model.User) (*model.User, error) {
	result := r.db.Model(&user).Association("Genres").Append(user.Genres)
	if result != nil {
		return nil, result
	}
	return user, nil
}
