package main

import "fmt"

func main() {
	x := 10

	if x > 5 {
		fmt.Println("x is greater than 5")
	} else if x < 5 && x != 0 {
		fmt.Println("x is less than 5")
	} else {
		fmt.Println("x is zero")
	}

	//---------------------------------------------------------------------

	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Unknown day")
	}

	// Example 2: Switch statement with multiple values
	month := "January"
	switch month {
	case "January", "February", "March":
		fmt.Println("Winter")
	case "April", "May", "June":
		fmt.Println("Spring")
	default:
		fmt.Println("Other season")
	}
	// Example 3: Switch with expression
	temperature := 25
	switch {
	case temperature < 0:
		fmt.Println("Freezing")
	case temperature >= 0 && temperature < 10:
		fmt.Println("Cold")
	case temperature >= 10 && temperature < 20:
		fmt.Println("Cool")
	case temperature >= 20 && temperature < 30:
		fmt.Println("Warm")
	default:
		fmt.Println("Hot")
	}

}
