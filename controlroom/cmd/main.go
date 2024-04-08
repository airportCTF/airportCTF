package main

import (
	"github.com/chessnok/airportCTF/controlroom/internal/application"
	"os"
)

// todo: создание рейсов, самолётов

// todo: доступ к этой функциональности только администратору (ходить в сервис авторизации и проверять JWT токен, либо, самостоятельно проверять JWT токен.)

func main() {
	os.Exit(MainWithCode())
}

func MainWithCode() int {
	app := application.NewApplication()
	app.Run()
	return 0
}
