package main

import "fmt"

func hello() {
	fmt.Println("Hello World !!!")
}

// func sum (a int, b int) int {}

// func sum(a, b int) int {
// 	return a + b
// }

func sum(a, b int) (result int) {
	result = a + b
	return
}

func main() {
	hello()

	ans := sum(6, 9)
	fmt.Println("Sum is : ", ans)
}
