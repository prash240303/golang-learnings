package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://portfolio-prash.vercel.app/"

func main() {
	fmt.Print("web request")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("response is of type %T\n", response)
	defer response.Body.Close() //callers responsibility to close the connection

	data, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	content := string(data)
	fmt.Printf(content)

}
