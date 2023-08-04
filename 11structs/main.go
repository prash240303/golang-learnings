package main

import "fmt"

func main() {
	fmt.Print("structs in golang")
	//on inherince in golang : no super or parent

	prashant := User{"prashant", "prash@go.dev", true, 24}
	fmt.Println(prashant)
	fmt.Printf("prash details are : %+v\n", prashant)
	fmt.Printf("name is %v\n", prashant.Name)

}

// type Info interface {
// 	Name() string
// }

type User struct {
	Name   string
	Email  string
	status bool
	Age    int
}
