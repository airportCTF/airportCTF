package http

import (
	"fmt"
	"github.com/chessnok/airportCTF/core/pkg/db"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strings"
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
	logginMiddleware := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			LoggedInHeader := c.Request().Header.Get("X-Data-Auth")
			fmt.Println("LoggedInHeader:", LoggedInHeader)
			LoggedInHeader = strings.ToLower(LoggedInHeader)
			if LoggedInHeader != "true" {
				fmt.Println("Header is not true, setting user to nil and isAdmin to false")
				c.Set("user", nil)
				c.Set("isAdmin", false)
				return next(c)
			}

			dataLogin := c.Request().Header.Get("X-Data-Login")
			fmt.Println("dataLogin:", dataLogin)
			if dataLogin == "" {
				fmt.Println("dataLogin is empty")
				logger.Println("dataLogin empty")
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{})
			}

			fmt.Println("Before calling GetFromDB")
			user, err := db.Users.GetFromDB(dataLogin)
			if err != nil {
				fmt.Println("Error from GetFromDB:", err)
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"error": "Internal Server Error",
				})
			}

			if user == nil {
				fmt.Println("User not found in database")
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{})
			}

			fmt.Println("User retrieved from database:", user)
			c.Set("user", user)
			c.Set("isAdmin", user.IsAdmin)
			return next(c)
		}
	}

	server.Use(logginMiddleware)
	g := server.Group("/v1")
	ag := g.Group("/admin")
	ag.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			isAdmin := c.Get("isAdmin")
			fmt.Println("isAdmin:", isAdmin)
			if isAdmin == false {
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

// NewFlight и DeleteFlight функции нужно определить в соответствии с вашим приложением
