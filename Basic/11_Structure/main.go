package main

import (
	"fmt"
	"math"
)

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

/************************ Attach methods to struct ************************/

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func (v Vertex) Abs() float64 {
	return math.Sqrt(float64(v.X)*float64(v.X) + float64(v.Y)*float64(v.Y))
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}

//*********************************** Struct Embedding *****************************************
/*
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
*/
