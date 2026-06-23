package main

import (
	"fmt"
	"math"
)

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

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hello()

	ans := sum(6, 9)
	fmt.Println("Sum is : ", ans)

	//------------------------------------
	// since functions
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

}
