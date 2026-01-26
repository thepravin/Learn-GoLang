package main

import (
	"fmt"
	"log"
	"os"
)

/*
- use error for expected failure states. If a function can fail during normal operation (e.g., file not found, network timeout, invalid user input), it should return an error.
- We shouldn't crash the program; we should just tell the user.
*/
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("denominator must not be zero")
	}

	return a / b, nil
}

func main1() {
	// ans, _ := divide(10, 0)
	// fmt.Println("Division is : ", ans)

	ans, err := divide(10, 0)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Division is : ", ans)

}

/*
# panic :

- stops the ordinary flow of control. When a function panics, its execution stops, and the program begins to crash, executing any deferred functions along the way.
- This is a bug or a critical system failure. Stop everything.
*/

func initDatabase() {
	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		// Critical failure: We cannot run without a DB
		panic("DB_URL environment variable is not set!")
	}
	// Connect to DB...
}

func main2() {
	initDatabase()
	// The code below will never run if initDatabase panics
	println("Database connected.")
}

/*
# Recover :

- is a built-in function that regains control of a panicking goroutine. It is strictly used inside defer. If you do not call recover, a panic will crash the entire program.

e.g: Imagine a web server where a specific handler has a bug (a division by zero). You don't want the whole website to go down for everyone just because one page is broken.

*/

func safeHandler() {
	// 1. Defer a function that contains recover
	defer func() {
		if r := recover(); r != nil {
			// 2. If we are here, a panic happened. Log it.
			log.Println("Recovered from panic:", r)
			fmt.Println("Sent '500 Internal Server Error' to user")
		}
	}()

	// 3. Simulate a crash (Divide by zero)
	performRiskyOperation()

	fmt.Println("Request processed successfully") // This won't print
}

func performRiskyOperation() {
	var a, b int = 10, 0
	result := a / b // PANIC: integer divide by zero
	fmt.Println(result)
}

func main() {
	fmt.Println("Server started...")
	safeHandler()
	fmt.Println("Server is still running! It didn't crash.")
}
