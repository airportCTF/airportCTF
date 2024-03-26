package application

import (
	"github.com/chessnok/airportCTF/core/pkg/db"
	"github.com/labstack/echo/v4"
	"log"
)

type Application struct {
	Server *echo.Echo
	DB     *db.Postgres
	Logger *log.Logger
}
