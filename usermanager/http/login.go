package http

import (
	"github.com/chessnok/airportCTF/core/pkg/db"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type LoginRequest struct {
	Login    string `json:"login" xml:"login"`       // information about login of user
	Password string `json:"password" xml:"password"` // unhashed password
}

func Login(db *db.Postgres, secret string) func(c echo.Context) error {
	return func(c echo.Context) error {
		req := new(LoginRequest)
		r := map[string]string{}
		if err := c.Bind(req); err != nil {

			r["status"] = err.Error()
			return c.JSON(http.StatusBadRequest, r)
		}
		u, err := db.Users.GetFromDB(req.Login)
		if err != nil {
			r["status"] = err.Error()
			return c.JSON(http.StatusInternalServerError, r)
		}
		if u == nil {
			r["status"] = "user not found"
			return c.JSON(http.StatusNotFound, r)
		}
		if !u.ComparePassword(req.Password) {
			r["status"] = "wrong password"
			return c.JSON(http.StatusForbidden, r)
		}
		now := time.Now()
		exp := now.Add(30 * 24 * time.Hour)
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"login": u.Login,
			"exp":   exp.Unix(),
		})
		tokenString, err := token.SignedString([]byte(secret))
		if err != nil {
			r["status"] = "error while generating jwt"
			c.JSON(http.StatusInternalServerError, r)
		}
		cookie := &http.Cookie{
			Name:     "session",
			Value:    tokenString,
			HttpOnly: true,
			Expires:  now.Add(30 * 24 * time.Hour),
			Path:     "/",
		}
		c.SetCookie(cookie)
		r["status"] = "ok"
		return c.JSON(http.StatusOK, r)
	}
}
