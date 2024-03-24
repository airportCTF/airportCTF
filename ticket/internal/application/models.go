package application

import (
	"github.com/chessnok/airportCTF/core/pkg/db"
	"github.com/labstack/echo/v4"
	"log/slog"
)

type Application struct {
	Logger *slog.Logger
	Server *echo.Echo
	DB     *db.Postgres
}
