package handler

import (
	"server/repository"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Handler struct {
	repo *repository.Repository
}

func NewHandler(repo *repository.Repository) *Handler {
	return &Handler{repo: repo}
}

func SetupRoutes(e *echo.Echo, db *gorm.DB) {

	repo := repository.NewRepository(db)
	h := NewHandler(repo)

	v1 := e.Group("/v1")
	{
		u := v1.Group("/users")
		{
			u.GET("", h.GetUsers)
			u.GET("/:id", h.GetUser)
			u.POST("", h.CreateUser)
			u.PUT("/:id", h.UpdateUser)
		}
		g := v1.Group("/genres")
		{
			g.GET("", h.GetGenres)
			g.GET("/:id", h.GetGenre)
			g.POST("", h.CreateGenre)
			g.PUT("/:id", h.UpdateGenre)
		}
		t := v1.Group("/tickets")
		{
			t.GET("/", h.GetTickets)
			t.GET(":id", h.GetTicket)
			t.POST("", h.CreateTicket)
			t.PUT("/:id", h.UpdateTicket)
		}
	}
}
