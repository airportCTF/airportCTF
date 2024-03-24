package application

import (
	"github.com/chessnok/airportCTF/core/pkg/db"
	http2 "github.com/chessnok/airportCTF/ticket/http"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
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
	a.Server.POST("/v1/tickets", http2.NewTicket)
	a.Server.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "pong")
	})
	a.Logger.Info("Starting Echo server", a.Server.Start("0.0.0.0:8010"))
}

func (a *Application) setupDB() {
	a.DB = db.NewPostgres(db.NewConfigFromEnv())
}
