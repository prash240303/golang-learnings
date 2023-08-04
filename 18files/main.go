package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Print("welcome to files in golang")
	content := "this needs to go in a file-learnefalf"

	file, err := os.Create("./myfile.txt")
	// if err != nil {
	// 	panic(err)
	// }
	handleErr(err)

	length, err := io.WriteString(file, content)
	handleErr(err)

	fmt.Println("lenght is ", length)
	file.Close()
	readfile("./myfile.txt")
}

func readfile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	handleErr(err)
	fmt.Println("text content inside the file is", string(databyte))
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
