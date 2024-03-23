package application

import (
	tHTTP "github.com/chessnok/airportCTF/ticket/http"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
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
	a.Server.POST("/v1/tickets", tHTTP.NewTicket)
	a.Server.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "pong")
	})
	a.Logger.Info("Starting Echo server", a.Server.Start("0.0.0.0:8010"))
}
