package application

import (
	"fmt"
	"os"
	"time"
)

type Application struct {
}

func (a *Application) Run() {
	for {
		data, err := os.ReadFile("art.txt")
		if err != nil {
			fmt.Println("Failed to print cool art")
			time.Sleep(5 * time.Second)
			continue
		}
		fmt.Println(string(data))
		time.Sleep(5 * time.Second)
	}
}
func NewApplication() *Application {
	return &Application{}
}
