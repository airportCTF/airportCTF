package http

import (
	"github.com/chessnok/airportCTF/core/pkg/db"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Login(db *db.Postgres) func(c echo.Context) error {
	return func(c echo.Context) error {
		// todo compare login and password hash with database and if it is true than give to client a JWT token
		return c.JSON(http.StatusCreated, u)
	}
}
