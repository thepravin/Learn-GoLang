package main

import "fmt"

func main() {

	fmt.Println("Start of the program")

	// The function inside defer will be executed when the surrounding function (main in this case) exits

	defer fmt.Println("This will be executed at the end")

	fmt.Println("Middle of the program")

	//--------------------------------------------------------

	// When you have multiple defer statements in a function, they are executed in a last-in, first-out (LIFO) order

	fmt.Println("Start of the program")
	defer fmt.Println("This will be executed second")
	defer fmt.Println("This will be executed first")
	fmt.Println("Middle of the program")
}
