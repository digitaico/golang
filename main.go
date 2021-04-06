package main

import (
	"fmt"
	"io/ioutil"
	"os"
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
	Age  int
}

func (e *Empleado) UpdateName(newName string) {
	e.Name = newName
}

func (e *Empleado) PrintName() {
	fmt.Println(e.Name)
}

type Hirer interface {
	Hired()
	Fired()
}

func (e Empleado) Hired() {
	fmt.Printf("el empleado %s fue contratdo", e.Name)
}

func (e Empleado) Fired() {
	fmt.Printf("el empleado %s fue despedido a los %d años de edad", e.Name, e.Age)
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
	empleado.Age = 45
	empleado.PrintName()
	empleado.Hired()
	empleado.Fired()

	miTexto := []byte("esto es lo que quiero escribir,  jea,  mi sonia preciosa hermosa")

	erro := ioutil.WriteFile("texto2.txt", miTexto, 0777)

	if erro != nil {
		fmt.Println(erro)
	}

	data, err := ioutil.ReadFile("texto2.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(string(data))

	f, err := os.OpenFile("texto.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString("textop adicional no existente previamente! jea\nMonocuco\nMarimonda"); err != nil {
		panic(err)
	}

	data, err = ioutil.ReadFile("texto.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))

}
