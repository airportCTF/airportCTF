package application

import (
	"fmt"
	"github.com/chessnok/airportCTF/core/pkg/db"
	http2 "github.com/chessnok/airportCTF/ticket/http"
	"log"
	"sync"
)

func (a *Application) Run() {
	a.Logger.Println("Starting application ticket")
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		a.Logger.Fatal(fmt.Sprintf("Error while running server, error: %v", a.Server.Start(":8010")))
		wg.Done()
	}()
	err := a.DB.Connect()
	if err != nil {
		a.Logger.Fatal(fmt.Sprintf("Error while connecting to DB, error: %v", err))
		wg.Done()
	}
	wg.Wait()
	return
}

func NewApplication() *Application {
	app := &Application{}
	app.setupLogger()
	app.setupDB()
	app.setupServer()
	return app
}
func (a *Application) setupLogger() {
	a.Logger = log.Default()
}

func (a *Application) setupServer() {
	a.Server = http2.NewServer(a.Logger, a.DB)
}

func (a *Application) setupDB() {
	a.DB = db.NewPostgres(db.NewConfigFromEnv())
}
