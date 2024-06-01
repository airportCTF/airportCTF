package application

import (
	"fmt"
	http2 "github.com/chessnok/airportCTF/controlroom/http"
	"github.com/chessnok/airportCTF/core/pkg/db"
	"log"
	"os"
	"sync"
	"time"
)

func (a *Application) Run() {
	go printFunnyASCIISticker()
	a.Logger.Println("Starting application control room")
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		a.Logger.Fatal(fmt.Sprintf("Error while running server, error: %v", a.Server.Start(":8012")))
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

// funny?
func printFunnyASCIISticker() {
	for {
		data, err := os.ReadFile("art.txt")
		if err != nil {
			fmt.Println("Failed to print cool art")
			time.Sleep(5 * time.Minute)
			continue
		}
		fmt.Println(string(data))
		time.Sleep(5 * time.Minute)
	}
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
