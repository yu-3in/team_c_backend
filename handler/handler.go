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

	e.GET("/", helloWorld)

	v1 := e.Group("/v1")
	{
		v1.POST("/login", h.Login)
		v1.POST("/signup", h.CreateUser)
		v1.POST("/logout", h.Logout)

		auth := v1.Group("", middleware.AuthenticationMiddleware)
		{
			auth.GET("/me", h.GetME)
			auth.PUT("/me", h.UpdateME)
			v1.DELETE("/me", h.DeleteME)
			auth.POST("/logout", h.Logout)

			u := auth.Group("/users")
			{
				u.GET("", h.GetUsers)
				u.GET("/:id", h.GetUser)
				u.POST("/genres", h.CreateUserGenre)

			}
			g := auth.Group("/genres")
			{
				g.GET("", h.GetGenres)
				g.GET("/:id", h.GetGenre)
				g.POST("", h.CreateGenre)
				g.PUT("/:id", h.UpdateGenre)
				g.DELETE("/:id", h.DeleteGenre)
			}
			t := auth.Group("/tickets")
			{
				t.GET("", h.GetTickets)
				t.GET("/:id", h.GetTicket)
				t.POST("", h.CreateTicket)
				t.PUT("/:id", h.UpdateTicket)
				t.DELETE("/:id", h.DeleteTicket)
			}
		}
	}
}

func helloWorld(c echo.Context) error {
	return c.JSON(200, "Hello World")
}
