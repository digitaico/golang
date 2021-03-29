package main

import (
	"fmt"
)

// App - struct que contiene dependencias, pointers a DB conection.
type App struct{}

// Run - sets up App
func (app *App) Run() error {
	fmt.Println("Setting App")
	handler := transportHTTP.nNewHandler()
	return nil
}

func main() {
	fmt.Println("Rest api go - jea")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error inicializando la App ")
		fmt.Println(err)
	}
}
