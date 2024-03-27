package http

import (
	"github.com/chessnok/airportCTF/core/pkg/db"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"log"
)

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

func NewServer(logger *log.Logger, db *db.Postgres) *echo.Echo {
	server := echo.New()
	server.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte("secret"), // todo get secret from env
	}))
	g := server.Group("/v1")
	//g.GET("/profile")                 // todo get profile info
	g.POST("/register", Register(db)) // todo register new user
	//g.GET("/login")                   // todo login through
	return server
}
