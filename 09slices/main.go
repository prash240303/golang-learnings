package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{2, 1, 3, 5, 4}
	fmt.Println(numbers)    // Output: [1 2 3 4 5]
	fmt.Println(numbers[1]) // Output: [1 2 3 4 5]
	// Slicing the original slice to create a new slice
	newSlice := numbers[0:3]

	fmt.Println(newSlice)

	//append slice elements
	numbers = append(numbers, 4, 5)
	fmt.Println(numbers)

	length := len(numbers)
	capacity := cap(numbers)

	fmt.Println("Length:", length) // Output: Length: 3
	fmt.Println("Capacity:", capacity)

	fmt.Println(sort.IntsAreSorted(numbers))
	sort.Ints(numbers)
	fmt.Println(numbers)
}
