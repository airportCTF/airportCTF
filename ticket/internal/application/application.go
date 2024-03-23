package application

import (
	"github.com/chessnok/airportCTF/core/pkg/db"
	"github.com/labstack/echo/v4"
	"log/slog"
	"os"
)

func (a *Application) Run() {
	a.Logger.Info("Starting application")
	a.Server.Logger.Fatal(a.Server.Start(":8080"))
}

func NewApplication() *Application {
	app := &Application{}
	app.setupLogger()
	app.setupServer()
	return app
}
func (a *Application) setupLogger() {
	a.Logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{}))
}

func (a *Application) setupServer() {
	a.Server = echo.New()
}

func (a *Application) setupDB() {
	a.DB = db.NewPostgres(db.NewConfigFromEnv())
}
