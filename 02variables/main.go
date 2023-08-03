package main

import (
	"fmt"
)

// //not allowed outside the method
// jwtToken:=3999
//var jwtToeken is allowrd

const LoginToken string = "thisistoken" //public

func main() {
	fmt.Println("variables")

	var username string = "hello"
	fmt.Printf(username)
	fmt.Printf("variable is of type %T \n", username)

	var isLooggedin bool = false
	fmt.Println(isLooggedin)
	fmt.Printf("variable is of type %T \n", isLooggedin)

	var smallval uint8 = 255
	var smallval2 uint8 = 254
	//unit 8 0-255
	var integer int = 25544

	fmt.Println(smallval)
	fmt.Println(smallval2)
	fmt.Println(integer)

	var website = "haha"
	fmt.Println(website)

	// not allowed
	// website = 112
	// fmt.Println(website)

	//no var style
	numberOfUsers := 3000
	numberOfUsers2 := "hahah"
	fmt.Println(numberOfUsers)
	fmt.Println(numberOfUsers2)

}
