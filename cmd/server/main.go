package main

import "fmt"

// App - struct que contiene dependencias, pointers a DB conection.
type App struct{}

func main() {
	fmt.Println("Rest api go - jea")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error inicializando la App ")
	}
}
