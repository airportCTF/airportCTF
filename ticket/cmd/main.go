package main

import (
	"os"

	"github.com/chessnok/airportCTF/ticket/internal/application"
)

func main() {
	os.Exit(MainWithCode())
}

func MainWithCode() int {
	app := application.NewApplication()
	app.Run()
	return 0
}
