package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// fmt.Printf("Enter number : ")
	// var num int

	// fmt.Scan(&num)
	// fmt.Println("Number is : ", num)

	// ----

	fmt.Printf("New Enter number : ")
	reader := bufio.NewReader(os.Stdin)
	newNum, _ := reader.ReadString('\n')
	fmt.Println("New Number is : ", newNum)
	fmt.Printf("Type of newNum : %T\n", newNum)
}
