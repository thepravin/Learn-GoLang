package main

import "fmt"

/*
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {

	// Method 1 : using var keyword
	var person1 Person
	person1.Age = 22
	person1.FirstName = "Pravin"
	person1.LastName = "Nalawade"

	fmt.Println(person1)

	//--------------------------------------------------
	// Method 2: Using a struct literal
	person3 := Person{
		FirstName: "Bob",
		LastName:  "Johnson",
		Age:       35,
	}

	fmt.Println(person3)

	//-------------------------------------------------------------
	// Method 3: Using the new keyword (returns a pointer to the struct)
	person4 := new(Person)
	person4.Age = 23
	person4.FirstName = "Pravin"
	person4.LastName = "Nalawade"
	fmt.Println(person4)
}

*/

//*********************************** Struct Embedding *****************************************

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Contact struct {
	Email string
	Phone string
}

type Address struct {
	Street  string
	City    string
	Country string
}

type Employee struct {
	Person   // Embedded struct
	Address  // Embedded struct
	Contact  // Embedded struct
	Position string
}

func main() {
	employee := Employee{
		Person: Person{
			FirstName: "Frank",
			LastName:  "Miller",
			Age:       45,
		},
		Address: Address{
			Street:  "123 Main St",
			City:    "Anytown",
			Country: "USA",
		},
		Contact: Contact{
			Email: "frank@example.com",
			Phone: "555-1234",
		},
		Position: "Manager",
	}

	fmt.Println(employee)
}
