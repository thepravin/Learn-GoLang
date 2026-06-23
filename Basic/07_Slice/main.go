/*
	- fixed and dynamic data structure
	- length can be changed during program's execution
	- Slices hold references to an underlying array, and if you assign one slice to another, both refer to the same array.

	- https://golangforall.com/en/post/golang-slice.html
*/

package main

import (
	"fmt"
)

func main() {
	// Slice deceleration
	q := []int{2, 3, 4, 5, 6}
	fmt.Println(q)

	r := []bool{true, false, true, false, false}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{4, true},
		{5, false},
		{6, false},
	}
	fmt.Println(s)

	slice1 := q[1:]
	slice2 := q[:5]
	slice3 := q[1:3]
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)

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
