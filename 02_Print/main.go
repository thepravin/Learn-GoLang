package main

import "fmt"

func main() {
	age := 22
	name := "Pravin"

	fmt.Println("Age : ", age, "Name : ", name)
	fmt.Println("Next line text")

	fmt.Printf("My age is %d\n", age)
	fmt.Printf("My name is %s\n", name)
	fmt.Printf("May age is %d and Name is %s\n", age, name)

	// If \n not added then cursor not shift to next line it remains same

	fmt.Printf("Type of name is : %T\n", name)

}
