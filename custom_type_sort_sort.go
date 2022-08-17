package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

// implement sort.Interface: https://cs.opensource.google/go/go/+/refs/tags/go1.19:src/sort/sort.go;l=14
type Persons []Person

func (p Persons) Len() int { return len(p) }
func (p Persons) Less(i, j int) bool {
	return p[i].Age < p[j].Age
}
func (p Persons) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func main() {
	persons := []Person{
		{"Cherry", 7},
		{"Alice", 22},
		{"Terry", 2},
		{"Jimmy", 18},
		{"Kenny", 30},
	}

	// use function sort.Sort
	personTypeSlice := persons
	sort.Sort(Persons(personTypeSlice))
	fmt.Println("Sorted persons based on Age by implementing sort interface: ", personTypeSlice)

	/* output
	Sorted persons based on Age by implementing sort interface:  [{Terry 2} {Cherry 7} {Jimmy 18} {Alice 22} {Kenny 30}]
	*/
}
