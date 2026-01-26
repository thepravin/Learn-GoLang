package main

import "fmt"

func main() {
	// Creating a map
	studentGrades := make(map[string]int)

	// Adding key-value pairs
	studentGrades["Alice"] = 90
	studentGrades["Bob"] = 85
	studentGrades["Charlie"] = 95

	// Accessing values
	fmt.Println("Alice's Grade:", studentGrades["Alice"])

	// Modifying values
	studentGrades["Bob"] = 88

	// Deleting a key-value pair
	delete(studentGrades, "Charlie")

	// Checking if a key exists
	grade, exists := studentGrades["David"]
	fmt.Println("David's Grade Exists:", exists)
	fmt.Println("David's Grade:", grade)

	// Iterating over the map
	fmt.Println("Student Grades:")
	for name, grade := range studentGrades {
		fmt.Printf("%s: %d\n", name, grade)
	}

	//--------------------------------------------------------------

	// Creating a map using literal

	studentGrades2 := map[string]int{
		"Pravin": 99,
		"Sachit": 100,
	}

	fmt.Println(studentGrades2)
}
