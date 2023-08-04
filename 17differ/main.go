package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("hhahaha") //will be execute at the end of funct
	fmt.Print("Hello, world!")

	//lifo execution of differ

	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")

	myDefer()

}

//hellow world !!

//stack 0 , 1 ,2 ,3 ,4 , 5
//hello world
//5, 4, 3 , 2,1, 0
//three two one hahaha

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
