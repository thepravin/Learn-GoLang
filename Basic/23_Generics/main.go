package main

import "fmt"

// return the index of x in s, or -1 if not found.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	si := []int{18, 13, -8, 0}
	fmt.Println(Index(si, 12))

	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "bar"))
	// fmt.Println(Index(ss, 2))

}
