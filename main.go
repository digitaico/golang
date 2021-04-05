package main

import (
	"fmt"
)

func myfunction(firstName string, apellido string) (string, error) {
	return firstName + " " + apellido, nil
}

func addOne() func() int {
	var x int

	return func() int {
		x++
		return x + 1
	}
}

type Empleado struct {
	Name string
}

func (e *Empleado) UpdateName(newName string) {
	e.Name = newName
}

func (e *Empleado) PrintName() {
	fmt.Println(e.Name)
}

func main() {
	days := [...]string{"Dom", "Lun", "Mar", "Mier", "Jue", "Vier", "Sab"}
	weekdays := days[1:6]

	fmt.Println(days[0])
	fmt.Println(days[5])
	fmt.Println(weekdays)

	yts := map[string]int{
		"platzi":  1000000,
		"r&u":     3000,
		"macardi": 15,
	}

	fmt.Println(yts["platzi"])

	type Person struct {
		name string
		age  int
	}

	jorge := Person{name: "Jorge Eduardo Ardila", age: 51}

	fmt.Println(jorge.age)

	type Team struct {
		name    string
		players [2]Person
	}

	var myTeam Team
	fmt.Println(myTeam)

	players := [...]Person{Person{name: "Forrest"}, Person{name: "Gordon"}}
	celtic := Team{name: "Celtic FC", players: players}

	fmt.Println(celtic)

	fullName, err := myfunction("elliot", "alderson")

	if err != nil {
		fmt.Println("Handler error case")
	}
	fmt.Println(fullName)

	myFunc := addOne()

	fmt.Println(myFunc())
	fmt.Println(myFunc())
	fmt.Println(myFunc())
	fmt.Println(myFunc())

	var empleado Empleado
	empleado.Name = "Jorge Eduardo Ardfila Muñoz"
	empleado.PrintName()
}
