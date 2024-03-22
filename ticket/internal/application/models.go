package application

import (
	"github.com/labstack/echo/v4"
	"log/slog"
)

type Application struct {
	Logger *slog.Logger
	Server *echo.Echo
	DB     string
}
