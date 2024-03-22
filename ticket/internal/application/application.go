package application

import (
	"github.com/labstack/echo/v4"
	"log/slog"
	"os"
)

func (a *Application) Run() {
	a.setupLogger()
	a.setupServer()

}

func (a *Application) setupLogger() {
	a.Logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{}))
}

func (a *Application) setupServer() {
	a.Server = echo.New()
}
