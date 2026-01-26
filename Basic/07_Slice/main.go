/*
	- fixed and dynamic data structure
	- length can be changed during program's execution
	-
*/

package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	fmt.Println("Slice : ", numbers)

	fmt.Println("Length of slice : ", len(numbers))
	fmt.Println("Capacity : ", cap(numbers))

	//--------------------------------------------------------------------

	// make(datatype, len , capacity)
	numbers2 := make([]int, 3, 5)

	fmt.Println("Slice : ", numbers2)
	fmt.Println("Lenght : ", len(numbers2))
	fmt.Println("Capacity : ", cap(numbers2))

	numbers2[0] = 10
	numbers2[1] = 11
	numbers2[2] = 12

	numbers2 = append(numbers2, 1)
	numbers2 = append(numbers2, 2)
	numbers2 = append(numbers2, 3)

	fmt.Println("Number2 : ", numbers2)

}
