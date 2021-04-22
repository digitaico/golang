package main

import "fmt"

func sum(a, b int) int {
	var x int = a
	var y int = b
	return x + y
}
func main() {
	fmt.Println(sum(5, 3))
}
