package http

import (
	user2 "github.com/chessnok/airportCTF/core/pkg/user"
	"github.com/labstack/echo/v4"
)

func GetProfile() func(c echo.Context) error {
	return func(c echo.Context) error {
		if c.Get("user") == nil {
			return c.JSON(401, map[string]string{"status": "not authorized"})
		}
		u := c.Get("user").(*user2.User)
		return c.JSON(200, u)
	}
}
