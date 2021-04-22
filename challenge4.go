package main

import "fmt"

func bigSum(ar []int64) int64 {
	var sum int64

	for i := int64(0); i < int64(len(ar)); i++ {
		sum += ar[i]
	}
	return sum
}

func main() {
	ar := []int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}
	fmt.Println(bigSum(ar))
}
