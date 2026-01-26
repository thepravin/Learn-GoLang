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

	// A classic mistake is trying to modify the value variable v. The variable v is a copy of the data. Changing it does nothing to the original slice.

	nums := []int{10, 20}

	for _, v := range nums {
		v = v + 5 // This only changes the local copy 'v'
	}
	fmt.Println(nums) // [10 20] -> No change

	//-----

	for i := range nums {
		nums[i] = nums[i] + 5 // Access the array directly
	}
	fmt.Println(nums) // [15 25] -> Changed!

}
