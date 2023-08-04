package main

import "fmt"

func main() {
	fmt.Print("structs in golang")
	//on inherince in golang : no super or parent

	prashant := User{"prashant", "prash@go.dev", true, 24}
	fmt.Println(prashant)
	fmt.Printf("prash details are : %+v\n", prashant)
	fmt.Printf("name is %v\n", prashant.Name)
	prashant.NewEmail()
	prashant.GetStatus()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("is user active? ", u.Status)
}
func (u User) NewEmail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is ", u.Email)

}
