package http

import (
	"github.com/chessnok/airportCTF/core/pkg/db"
	"github.com/labstack/echo/v4"
)

func DeleteFlight(db *db.Postgres) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(200, "OK")
	}
}
