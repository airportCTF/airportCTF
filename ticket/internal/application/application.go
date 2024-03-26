package application

import (
	"fmt"
	"github.com/chessnok/airportCTF/core/pkg/db"
	http2 "github.com/chessnok/airportCTF/ticket/http"
	"log/slog"
	"os"
	"sync"
)

func (a *Application) Run() {
	a.Logger.Info("Starting application")
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		a.Logger.Error(fmt.Sprintf("Error while running server, error: %v", a.Server.Start(":8080")))
		wg.Done()
	}()
	err := a.DB.Connect()
	if err != nil {
		a.Logger.Info(fmt.Sprintf("Didn't connect to DB, error: %v", err))
		wg.Done()
	}
	wg.Wait()
	return
}

func NewApplication() *Application {
	app := &Application{}
	app.setupLogger()
	app.setupServer()
	app.setupDB()
	return app
}
func (a *Application) setupLogger() {
	a.Logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{}))
}

func (a *Application) setupServer() {
	a.Server = http2.NewServer()
}

func (a *Application) setupDB() {
	a.DB = db.NewPostgres(db.NewConfigFromEnv())
}
