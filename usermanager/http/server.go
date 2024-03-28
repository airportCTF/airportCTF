package http

import (
	"github.com/chessnok/airportCTF/core/pkg/db"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"log"
	"os"
)

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

func NewServer(logger *log.Logger, db *db.Postgres) *echo.Echo {
	server := echo.New()
	g := server.Group("/v1")
	//g.GET("/profile")                 // todo get profile info
	g.POST("/register", Register(db)) // todo register new user
	g.POST("/login", Login(db, os.Getenv("SECRET_KEY")))
	//g.GET("/login")                   // todo login through
	return server
}
