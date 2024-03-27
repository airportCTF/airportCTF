package http

import (
	"github.com/chessnok/airportCTF/core/pkg/db"
	"github.com/chessnok/airportCTF/core/pkg/user"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Register(db *db.Postgres) func(c echo.Context) error {
	return func(c echo.Context) error {
		u := new(user.User)
		if err := c.Bind(u); err != nil {
			return err
		}
		if err := db.Users.PutToDB(u); err != nil {
			return echo.ErrBadRequest
		}
		return c.JSON(http.StatusCreated, u)
	}
}
