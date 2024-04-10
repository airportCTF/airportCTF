package http

import (
	"github.com/chessnok/airportCTF/core/pkg/db"
	"github.com/chessnok/airportCTF/core/pkg/flight"
	"github.com/labstack/echo/v4"
)

func NewFlight(db *db.Postgres) echo.HandlerFunc {
	return func(c echo.Context) error {
		fl := flight.Flight{}
		if err := c.Bind(&fl); err != nil {
			return c.JSON(400, "Bad request")
		}
		if fl.To == "" || fl.From == "" || fl.Date.IsZero() || fl.ID == "" {
			return c.JSON(400, map[string]string{"error": "Bad request"})
		}
		err := db.Flights.PutToDB(&fl)
		if err != nil {
			return c.JSON(500, map[string]string{"error": "Internal server error"})
		}
		return c.JSON(200, fl)
	}
}
