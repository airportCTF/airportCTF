package http

import (
	"github.com/chessnok/airportCTF/core/pkg/db"
	"github.com/chessnok/airportCTF/core/pkg/ticket"
	"github.com/labstack/echo/v4"
	"net/http"
)

func NewTicket(db *db.Postgres) func(c echo.Context) error {
	return func(c echo.Context) error {
		t := new(ticket.Ticket)
		if err := c.Bind(t); err != nil {
			return err
		}
		if err := db.Tickets.PutToDB(t); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusCreated, t)
	}
}
