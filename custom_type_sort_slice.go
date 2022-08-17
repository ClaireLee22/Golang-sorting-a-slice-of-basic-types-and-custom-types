package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	persons := []Person{
		{"Cherry", 7},
		{"Alice", 22},
		{"Terry", 2},
		{"Jimmy", 18},
		{"Kenny", 30},
	}
	// use function sort.Slice: https://cs.opensource.google/go/go/+/refs/tags/go1.19:src/sort/slice.go;l=18
	sort.Slice(persons, func(i, j int) bool {
		return persons[i].Age < persons[j].Age
	})
	fmt.Println("Sorted persons based on Age: ", persons)

	sort.Slice(persons, func(i, j int) bool {
		return persons[i].Name < persons[j].Name
	})
	fmt.Println("Sorted persons based on Name: ", persons)

	/* output
	Sorted persons based on Age:  [{Terry 2} {Cherry 7} {Jimmy 18} {Alice 22} {Kenny 30}]
	Sorted persons based on Name:  [{Alice 22} {Cherry 7} {Jimmy 18} {Kenny 30} {Terry 2}]
	*/
}
