package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to the time study tut")
	presentTime := time.Now()
	currentTime := time.Now().Format("03:04 PM")

	// Get the current date in the "12-03-2003" format
	currentDate := time.Now().Format("01-02-2006")

	// Print the current time and date
	fmt.Printf("Current Time (12-hour format): %s\n", currentTime)
	fmt.Printf("Current Date (12-03-2003 format): %s\n", currentDate)
	fmt.Println(presentTime.Format("2023-08-03  2:54:09 Monday"))

	// preciseTime :=time.Now()

}
