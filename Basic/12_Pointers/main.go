package main

import "fmt"

func modifyValueByReference(num *int) {
	*num = *num * 2
}

func main() {

	// Method 1 :
	var num int = 2
	var ptr *int

	ptr = &num

	fmt.Println("Value of num : ", num)
	fmt.Println("Value through pointer : ", *ptr)

	// Method 2 :

	data := "pravin"
	pointer := &data
	fmt.Println("Value through pointer : ", *pointer)

	//--------------- nil pointer ----------------------

	var ptr2 *int

	if ptr2 == nil {
		fmt.Println("Pointer is nil")
	}

	//------------------ Pass by reference -------------------

	value := 10
	modifyValueByReference(&value)
	fmt.Println("Modified value : ", value)
}
