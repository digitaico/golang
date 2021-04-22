package main

import "fmt"

var ar = []int{1, 2, 3, 4, 10, 11}

func sumArray(ar []int) int {
	var sum int = 0
	for i := 0; i < int(len(ar)); i++ {
		sum += ar[i]
	}
	return sum
}
func main() {
	fmt.Println(sumArray(ar))
}
