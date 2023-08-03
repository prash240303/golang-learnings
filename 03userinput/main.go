package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input program"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the rating of out restrau")

	//comma ok //err ok : golang dont ahve try catdch

	input, err := reader.ReadString('\n')
	fmt.Println("thanks for rating ", input)
	fmt.Printf("type of this  for rating is %T", input)

}
