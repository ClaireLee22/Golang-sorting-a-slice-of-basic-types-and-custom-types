package main

import (
	"fmt"
	"sort"
)

func main() {

	// ints
	numbers := []int{7, 22, 2, 18, 30}
	sort.Ints(numbers)
	fmt.Println("Sorted ints in ascending order: ", numbers)

	// float64s
	float64s := []float64{100.2, 22.2, 50.0, 1.2, 12.8}
	sort.Float64s(float64s)
	fmt.Println("Sorted float64s in ascending order: ", float64s)

	// strings
	strings := []string{"Cherry", "Alice", "Terry", "Jimmy", "Kenny"}
	sort.Strings(strings)
	fmt.Println("Sorted strings in ascending order: ", strings)

	/* output
	Sorted ints in ascending order:  [2 7 18 22 30]
	Sorted float64s in ascending order:  [1.2 12.8 22.2 50 100.2]
	Sorted strings in ascending order:  [Alice Cherry Jimmy Kenny Terry]
	*/
}
