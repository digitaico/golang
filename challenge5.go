package main

import (
	"fmt"
	"math"
)

func diagonalDiff(arr [][]int) float64 {
	var sum1, sum2 int = 0, 0

	for i := int(0); i < int(len(arr)); i++ {
		sum1 += arr[i][i]
		sum2 += arr[i][int(len(arr))-1-i]
	}
	return math.Abs(float64(sum1 - sum2))
}

func main() {
	arr := [][]int{}
	fmt.Println(diagonalDiff(arr))
}
