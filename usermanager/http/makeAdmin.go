package http

import (
	"github.com/chessnok/airportCTF/core/pkg/db"
	"github.com/labstack/echo/v4"
	"os"
)

func MakeAdmin(db *db.Postgres) echo.HandlerFunc {
	return func(c echo.Context) error {
		header := c.Request().Header.Get("Authorization")
		if "Bearer "+os.Getenv("API_KEY") != header {
			return c.JSON(403, map[string]string{"error": "Forbidden"})
		}
		user := c.QueryParam("user")
		if user == "" {
			return c.JSON(400, map[string]string{"error": "Bad request"})
		}
		err := db.Users.MakeAdmin(user)
		if err != nil {
			return c.JSON(500, map[string]string{"error": "Internal server error"})
		}
		return c.JSON(200, map[string]string{"status": "Made admin successfully, if user exists"})
	}
}
