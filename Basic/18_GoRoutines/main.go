/*
# What is a Goroutine?

- A goroutine is a lightweight thread of execution managed by the Go runtime, not the operating system.
- Think of it as a function that can run concurrently (or in parallel) with other functions.
- The main program itself runs in a special goroutine called the "main goroutine."

    - Concurrency: Dealing with many things at once.
    - Parallelism: Doing many things at once.

Goroutines are the building blocks of concurrency in Go.

- Creating a goroutine is simple: just add the go keyword before a function call.

- A Go program exits immediately when the main goroutine finishes, even if other goroutines are still running.
	- You must synchronize your goroutines to ensure main waits for them.

# How to Synchronize Goroutines?

There are two primary ways to manage and synchronize goroutines:

   -  sync.WaitGroup: When you just need to wait for a group of goroutines to finish.

    - Channels: When you need goroutines to communicate with each other.

*/

/*

package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 3; i++ {
		fmt.Println("Goroutine...")
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	// 1. Start a new goroutine
	go sayHello()

	// 2. The main goroutine continues...
	for i := 0; i < 3; i++ {
		fmt.Println("Main!")
		time.Sleep(100 * time.Millisecond)
	}
}

*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func doWrok(id int, wg *sync.WaitGroup) {

	// Tell the waitgroup we are done when the function return
	defer wg.Done()

	fmt.Printf("Worker %d starting.....\n", id)
	time.Sleep(1 * time.Second)
	fmt.Printf("Worker %d finished.\n", id)
}

func main() {
	// 1. Create a waitgroup
	var wg sync.WaitGroup

	//2. Launch 3 goroutines
	for i := 1; i <= 3; i++ {
		// 3. Increment the counter for each goroutine
		wg.Add(1)
		go doWrok(i, &wg)
	}

	fmt.Println("Main is waiting for workers to finish....")

	wg.Wait()

	fmt.Println("All workers finished...Main exiting.")
}
