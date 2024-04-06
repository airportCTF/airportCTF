package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func Logout() func(c echo.Context) error {
	return func(c echo.Context) error {
		cookie := &http.Cookie{
			Name:     "session",
			Value:    "",
			HttpOnly: true,
			Expires:  time.Now().Add(-1 * time.Hour),
		}
		c.SetCookie(cookie)
		red_url := c.QueryParam("redirect")
		if red_url == "" {
			red_url = "/"
		}
		return c.Redirect(http.StatusMovedPermanently, red_url)
	}
}
