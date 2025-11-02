/*
	- no do-while loop in go
	- no while loop in go
*/

package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 5; i++ {
		fmt.Print(i, " ")
	}

	//------------------------------------------------------

	// Example 2: Infinite loop with break statement
	counter := 0
	for {
		fmt.Println("Infinite Loop")
		counter++
		if counter == 3 {
			break
		}
	}

	//----------------------- Range ---------------------------------

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for index, value := range numbers {
		fmt.Printf("Index : %d, Value : %d\n", index, value)
	}

	message := "Hello, Golang!"
	for index, char := range message {
		fmt.Printf("Index: %d, Character: %c\n", index, char)
	}

}
