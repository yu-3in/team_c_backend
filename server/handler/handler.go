package handler

import (
	"server/middleware"
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
		v1.POST("/login", h.Login)
		v1.POST("/signup", h.CreateUser)
		v1.POST("/logout", h.Logout)

		auth := v1.Group("", middleware.AuthenticationMiddleware)
		{
			auth.GET("/me", h.GetME)
			auth.PUT("/me", h.UpdateME)

			u := auth.Group("/users")
			{
				u.GET("", h.GetUsers)
				u.GET("/:id", h.GetUser)

			}
			g := auth.Group("/genres")
			{
				g.GET("", h.GetGenres)
				g.GET("/:id", h.GetGenre)
				g.POST("", h.CreateGenre)
				g.PUT("/:id", h.UpdateGenre)
			}
			t := auth.Group("/tickets")
			{
				t.GET("", h.GetTickets)
				t.GET("/:id", h.GetTicket)
				t.POST("", h.CreateTicket)
				t.PUT("/:id", h.UpdateTicket)
			}
		}
	}
}
