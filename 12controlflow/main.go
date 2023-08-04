package main

import (
	"fmt"
)

func main() {
	fmt.Println("if else statemens")
	loginCount := 23

	if loginCount < 10 {
		fmt.Println("regular user")
	} else {
		fmt.Println("super user")
	}
}
