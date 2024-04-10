package http

import (
	"github.com/chessnok/airportCTF/core/pkg/db"
	"github.com/labstack/echo/v4"
)

type DeleteFlightRequest struct {
	ID string `json:"id"`
}

func DeleteFlight(db *db.Postgres) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := DeleteFlightRequest{}
		if err := c.Bind(&req); err != nil {
			return c.JSON(400, "Bad request")
		}
		if req.ID == "" {
			return c.JSON(400, map[string]string{"error": "Bad request"})
		}
		err := db.Flights.DeleteFromDB(req.ID)
		if err != nil {
			return c.JSON(500, map[string]string{"error": "Internal server error"})
		}
		return c.JSON(200, map[string]string{"status": "Deleted flight if exists"})
	}
}
