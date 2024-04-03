package http

import (
	"github.com/chessnok/airportCTF/core/pkg/db"
	"github.com/labstack/echo/v4"
	"log"
)

func NewLoggingMiddleware(logger *log.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			logger.Printf("Request: %s %s", c.Request().Method, c.Request().URL.Path)
			for k, v := range c.Request().Header {
				logger.Printf("Header %s: %s", k, v)
			}
			return next(c)
		}
	}
}
func NewServer(logger *log.Logger, db *db.Postgres) *echo.Echo {
	server := echo.New()
	server.Use(NewLoggingMiddleware(logger))
	g := server.Group("/v1")
	g.POST("/tickets", NewBooking(db, logger))
	g.GET("/tickets", GetTickets(db))
	return server
}
