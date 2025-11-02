package main

import (
	"fmt"
	"strings"
)

func main() {

	//------------------- Split -----------------------

	str := "apple,orange,banana"
	parts := strings.Split(str, ",")
	fmt.Println(parts) // Output: [apple orange banana]

	//--------------------- Count ------------------------

	str2 := "one two three four two two five"
	count := strings.Count(str2, "two")
	fmt.Println("Count:", count)

	//---------------------- TrimSpace --------------------

	str3 := " Hello, Go! "
	trimmed := strings.TrimSpace(str3)
	fmt.Println("Trimmed:", trimmed)

	//--------------------- Concatenating --------------------

	result := strings.Join([]string{str, str2}, " ")
	fmt.Println(result)

}
