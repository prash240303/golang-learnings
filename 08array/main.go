package main

import "fmt"

func main() {
	fmt.Println("wecome to the array in golang")

	var fruitList [4]string
	fruitList[0] = "apple"
	fruitList[1] = "banana"
	fruitList[2] = "30"
	fruitList[3] = "'40'"

	fmt.Println(fruitList[0])

	nums := [4]int{1, 2, 3, 4}
	fmt.Println(nums[2])
	fmt.Println("length of nums arr", len(nums))

}
