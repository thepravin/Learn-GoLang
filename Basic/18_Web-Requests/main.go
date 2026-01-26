package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Learning web request...")

	// Make a GET request
	response, err :=
		http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Println("Error making GET request: ", err)
		return
	}

	defer response.Body.Close()

	fmt.Printf("Type of response: %T\n", response)

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Print the response body as a string
	fmt.Println(string(body))
}
