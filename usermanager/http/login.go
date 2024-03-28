package http

import (
	"github.com/chessnok/airportCTF/core/pkg/db"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type LoginRequest struct {
	Login    string `json:"login" xml:"login"`
	Password string `json:"password" xml:"password"`
}

type LoginResponse struct {
	Token string `json:"Token" xml:"Token"`
}

func Login(db *db.Postgres, secret string) func(c echo.Context) error {
	return func(c echo.Context) error {
		req := new(LoginRequest)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		u, err := db.Users.GetFromDB(req.Login)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		if u == nil {
			return c.JSON(http.StatusNotFound, "user not found")
		}
		if !u.ComparePassword(req.Password) {
			return c.JSON(http.StatusForbidden, "wrong password")
		}
		now := time.Now()
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"login":    u.Login,
			"isAdmin":  u.IsAdmin,
			"lastName": u.LastName,
			"name":     u.Name,
			"exp":      now.Add(30 * 24 * time.Hour).Unix(),
		})
		tokenString, err := token.SignedString([]byte(secret))
		if err != nil {
			c.JSON(http.StatusInternalServerError, "error while generating jwt")
		}
		r := new(LoginResponse)
		r.Token = tokenString
		return c.JSON(http.StatusOK, r)
	}
}
