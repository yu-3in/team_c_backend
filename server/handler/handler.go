package handler

import (
	"server/repository"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"net/http"
	"time"
)

type Handler struct {
	repo *repository.Repository
}

func NewHandler(repo *repository.Repository) *Handler {
	return &Handler{repo: repo}
}

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	name := claims.Name
	return c.JSON(http.StatusOK, "Welcome "+name+"!")
}

func login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Throws unauthorized error
	if username != "jon" || password != "shhh!" {
		return echo.ErrUnauthorized
	}

	// Set custom claims
	claims := &jwtCustomClaims{
		"Jon Snow",
		true,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func SetupRoutes(e *echo.Echo, db *gorm.DB) {

	repo := repository.NewRepository(db)
	h := NewHandler(repo)

	e.POST("/login", login)
	r := e.Group("/restricted")
	{
		// Configure middleware with the custom claims type
		config := echojwt.Config{
			NewClaimsFunc: func(c echo.Context) jwt.Claims {
				return new(jwtCustomClaims)
			},
			SigningKey: []byte("secret"),
		}
		r.Use(echojwt.WithConfig(config))
		r.GET("", restricted)
	}

	v1 := e.Group("/v1/")
	{
		u := v1.Group("users/")
		{
			u.GET("", h.GetUsers)
			u.GET(":id", h.GetUser)
		}
		g := v1.Group("genres/")
		{
			g.GET("", h.GetGenres)
			g.GET(":id", h.GetGenre)
		}
		t := v1.Group("tickets/")
		{
			t.GET("", h.GetTickets)
			t.GET(":id", h.GetTicket)
		}
	}
}
