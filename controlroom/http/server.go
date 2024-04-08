package http

import (
	"github.com/chessnok/airportCTF/core/pkg/db"
	"github.com/labstack/echo/v4"
	"log"
	"strings"
)

func NewServer(logger *log.Logger, db *db.Postgres) *echo.Echo {
	server := echo.New()
	logginMiddlewar := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			LoggedInHeader := c.Request().Header.Get("X-Data-Auth")
			LoggedInHeader = strings.ToLower(LoggedInHeader)
			if LoggedInHeader != "true" {
				c.Set("user", nil)
				c.Set("isAdmin", false)
				return next(c)
			}
			user, _ := db.Users.GetFromDB(c.Request().Header.Get("X-Data-Login"))
			c.Set("user", user)
			c.Set("isAdmin", user.IsAdmin)
			return next(c)
		}
	}
	server.Use(logginMiddlewar)
	g := server.Group("/v1")
	ag := g.Group("/admin")
	ag.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if c.Get("isAdmin") == false {
				return c.JSON(403, map[string]string{"error": "Forbidden"})
			}
			return next(c)
		}
	})
	fg := ag.Group("/flight")
	fg.POST("/new", NewFlight(db))
	fg.DELETE("/delete", DeleteFlight(db))
	return server
}
