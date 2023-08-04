package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// rand.Seed()
	// NewRand(NewSource(time.Now().UnixNano()))

	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	diceNumber := rng.Intn(6) + 1 // Generate random integer between 0 and 99.
	fmt.Println("value of dice is", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("dice value is 1 and u can open")
	case 2:
		fmt.Println("u can move 2 spts")
		fallthrough
	case 3:
		fmt.Println("u can move 3 spts")
		fallthrough
	case 4:
		fmt.Println("u can move 4 spts")
		fallthrough
	case 5:
		fmt.Println("u can move 5 spts")
	case 6:
		fmt.Println("u can redraw")
	}
}
