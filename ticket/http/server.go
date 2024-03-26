package http

import (
	"github.com/chessnok/airportCTF/core/pkg/ticket"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func NewTicket(c echo.Context) error {
	t := new(ticket.Ticket)
	if err := c.Bind(t); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, t)
}

func NewLoggingMiddleware(logger *log.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			logger.Printf("Request: %s %s", c.Request().Method, c.Request().URL.Path)
			return next(c)
		}
	}
}
func NewServer(logger *log.Logger) *echo.Echo {
	server := echo.New()
	server.Use(NewLoggingMiddleware(logger))
	server.POST("/v1/tickets", NewTicket)
	server.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "pong")
	})
	return server
}
