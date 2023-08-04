package main

import "fmt"

func main() {
	// fmt.Print("Hello, world!")
	days := []string{"sunday", "monday", "wednesday", "friday"}
	// fmt.Print(days)
	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	//ragne loop
	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for index, day := range days {
		fmt.Printf("index is %v and value is %v\n", index, day)
	}

	rougeValue := 1
	for rougeValue < 10 {
		fmt.Printf("index is %v ", rougeValue)
		rougeValue++
	}
}
