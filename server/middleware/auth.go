package middleware

import (
	"server/util"

	"github.com/labstack/echo/v4"
)

func AuthenticationMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, err := util.ParseToken(c)
		if err != nil {
			return echo.ErrUnauthorized
		}
		c.Set("userID", userID)

		return next(c)
	}
}
