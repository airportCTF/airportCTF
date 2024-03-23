package http

import (
	"github.com/chessnok/airportCTF/core/pkg/ticket"
	"github.com/labstack/echo/v4"
	"net/http"
)

func NewTicket(c echo.Context) error {
	t := new(ticket.Ticket)
	if err := c.Bind(t); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, t)
}
