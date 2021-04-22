package main

// comparing triplets
import "fmt"

func compareTriplets(ar []int, br []int) []int {
	var result []int
	aSum := 0
	bSum := 0

	for i := 0; i < int(len(ar)); i++ {
		if ar[i] > br[i] {
			aSum++
		}
		if ar[i] < br[i] {
			bSum++
		}
	}

	result = append(result, aSum, bSum)
	return result
}

func main() {
	ar := []int{17, 28, 30}
	br := []int{99, 16, 8}

	fmt.Println(compareTriplets(ar, br))
}
