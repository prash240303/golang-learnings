//new() allocate memory but not initalize
//get a mem addess
//zeroed storage

// make() : allocate memory  and initalized
// you will get a memory add
//non zeroed storage :go ahead and put data

// gc :garbage collection happens automaticall

package main

import (
	"fmt"
	"runtime"
)

func main() {
	numCPU := runtime.NumCPU()
	numCGOCall := runtime.NumCgoCall()
	fmt.Println("Number of logical CPUs:", numCPU)
	fmt.Println("Number of cgo CPUs:", numCGOCall)
}
