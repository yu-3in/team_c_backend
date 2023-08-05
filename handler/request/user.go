package request

import (
	"server/model"
	"time"
)

type ReqUpdateUser struct {
	ID             int    `param:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	DepartmentName string `json:"departmentName"`
	ProductName    string `json:"productName"`
}

type ReqLoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ReqUpdateUserGenre struct {
	GenreID []int `json:"genreIds"`
}

type ReqCreateUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResUser struct {
	ID             int           `json:"id"`
	Name           string        `json:"name"`
	Email          string        `json:"email"`
	IconColor      string        `json:"iconColor"`
	DepartmentName string        `json:"departmentName"`
	ProductName    string        `json:"productName"`
	Genres         []model.Genre `json:"genres"`
	CreatedAt      time.Time     `json:"createdAt"`
	UpdatedAt      time.Time     `json:"updatedAt"`
}

func UserModelToResUser(m *model.User) *ResUser {
	return &ResUser{
		ID:             m.ID,
		Name:           m.Name,
		Email:          m.Email,
		IconColor:      m.IconColor,
		DepartmentName: m.DepartmentName,
		ProductName:    m.ProductName,
		Genres:         m.Genres,
		CreatedAt:      m.CreatedAt,
		UpdatedAt:      m.UpdatedAt,
	}
}

func UserModelToResUsers(m []*model.User) []*ResUser {
	responses := make([]*ResUser, len(m))
	for i, user := range m {
		responses[i] = UserModelToResUser(user)
	}
	return responses
}

// []*model.Userから[]model.Userに変換するヘルパー関数
func ConvertToModelUserSlice(users []model.User) []*model.User {
	result := make([]*model.User, len(users))
	for i, user := range users {
		result[i] = &user
	}
	return result
}
