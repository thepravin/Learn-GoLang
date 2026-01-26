package main

import (
	"fmt"
	"strconv"
)

func main() {

	// Numberic conversion :
	var integerNum int = 12
	var floatNum float64 = float64(integerNum)
	fmt.Println("Integer to Float : ", floatNum)

	// String conversion :
	str := strconv.Itoa(integerNum)
	fmt.Println("Num to String : ", str)

	strNum := "123"
	num, err := strconv.Atoi(strNum)
	if err == nil {
		fmt.Println("String to Num : ", num)
	}

	strFloat := "3.14"
	num2, err2 := strconv.ParseFloat(strFloat, 64)
	if err2 == nil {
		fmt.Println(num2)
	}
}
