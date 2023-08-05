package request

import (
	"server/model"
	"time"
)

type ReqCreateGenre struct {
	Title string `json:"title"`
	Color string `json:"color"`
}

type ReqUpdateGenre struct {
	ID    int    `param:"id"`
	Title string `json:"title"`
	Color string `json:"color"`
}

type ResGenre struct {
	ID        int        `json:"id" gorm:"primaryKey"`
	Title     string     `json:"title"`
	Users     []*ResUser `json:"users" gorm:"many2many:user_genres;"`
	Color     string     `json:"color"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
}

func GenreModelToResGenre(m *model.Genre) *ResGenre {
	return &ResGenre{
		ID:        m.ID,
		Title:     m.Title,
		Users:     UserModelToResUsers(ConvertToModelUserSlice(m.Users)),
		Color:     m.Color,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func GenreModelToResGenres(m []*model.Genre) []*ResGenre {
	responses := make([]*ResGenre, len(m))
	for i, genre := range m {
		responses[i] = GenreModelToResGenre(genre)
	}
	return responses
}
