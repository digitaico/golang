package main

// plusminus
import "fmt"

func plusMinus(arr []int) {
	var positivos, negativos, ceros int = 0, 0, 0
	l := int(len(arr))

	for i := int(0); i < l; i++ {
		if arr[i] > 0 {
			positivos++
		}
		if arr[i] < 0 {
			negativos++
		}
		if arr[i] == 0 {
			ceros++
		}
	}
	fmt.Printf("%.6f\n", float64(positivos)/float64(l))
	fmt.Printf("%.6f\n", float64(negativos)/float64(l))
	fmt.Printf("%.6f\n", float64(ceros)/float64(l))
}

func main() {
	arr := []int{-4, 3, -9, 0, 4, 1}
	plusMinus(arr)
}
