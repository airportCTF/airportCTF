package main

import (
	"github.com/chessnok/airportCTF/usermanager/internal/application"
	"os"
)

func main() {
	os.Exit(MainWithCode())
}

func MainWithCode() int {
	app := application.NewApplication()
	app.Run()
	return 0
}
