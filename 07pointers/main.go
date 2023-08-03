package main

import "fmt"

func main() {
	fmt.Println("welcome to a class on pointers")

	// var one int = 3
	//create a pointer
	// var ptr *int
	// fmt.Println("value of pointer is ", ptr)

	myNumber := 23
	var numPointer *int = &myNumber
	fmt.Println("value of pointer is (Address stored of mynimber) ", numPointer)
	fmt.Println("value pointed by the num is ", *numPointer)

}
