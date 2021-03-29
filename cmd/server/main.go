package main

import (
	"fmt"
	"net/http"

	"github.com/digitai/golang/go-rest-api/internal/database"
	transportHTTP "github.com/digitai/golang/go-rest-api/internal/transport/http"
)

// App - struct que contiene dependencias, pointers a DB conection.
type App struct{}

// Run - sets up App
func (app *App) Run() error {
	fmt.Println("Setting Up App")

	var err error
	_, err = database.NewDatabase()
	if err != nil {
		return err
	}

	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Fallo en setup se server")
		return err
	}

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
