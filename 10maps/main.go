package main

import "fmt"

func main() {
	fmt.Println("maps in golang")
	languages := make(map[string]string)

	languages["JS"] = "javascript"
	languages["RB"] = "ruby"
	languages["PY"] = "python"

	fmt.Println("list of all langs ", languages)
	fmt.Println("js shorts for ", languages["JS"])
	delete(languages, "PY")
	fmt.Println("list of all langs ", languages)

	//loops in golang

	for key, value := range languages {
		fmt.Printf("for key %v, value is %v", key, value)
	}
}
