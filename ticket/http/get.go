package http

import (
	"github.com/chessnok/airportCTF/core/pkg/db"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetTickets(db *db.Postgres) func(c echo.Context) error {
	return func(c echo.Context) error {
		tickets, err := db.Tickets.GetAllFromDB()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, tickets)
	}
}
