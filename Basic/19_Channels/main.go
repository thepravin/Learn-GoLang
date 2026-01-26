/*
	- Channels are medium for the communicating goroutines or sharing the data between them.
	- Channels work (FIFO) First in first out.

	# Types of channels :

*/

/*

	1) Unbuffered Channels (Synchronous) :
		- This is the default type. An unbuffered channel has zero capacity.

	How to create: ch := make(chan int) (No second argument)

	How it works: It forces perfect synchronization.

		A send operation (ch <- data) will block until another goroutine is ready to receive (<-ch).

		A receive operation (<-ch) will block until another goroutine sends (ch <- data).

	- This is like a direct handoff.

*/

/*
package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("Working...")
	time.Sleep(1 * time.Second)
	fmt.Println("Done")

	// Send a signal; this will block until main is ready to receive it
	done <- true
}

func main() {
	done := make(chan bool) // Create an unbuffered channel

	go worker(done)

	// Block the main goroutine until it receives a signal from the worker
	<-done

	fmt.Println("Main goroutine is finished.")

}
*/

/*
2. Buffered Channels (Asynchronous)

A buffered channel has a fixed capacity (a "buffer" size greater than zero).

    How to create: ch := make(chan string, 10) (The second argument is the capacity)

    How it works: It decouples the sender and receiver.

        A send operation (ch <- data) will only block if the buffer is full.

        A receive operation (<-ch) will only block if the buffer is empty.

This is like a mailbox. A sender can drop off several letters (up to the buffer capacity) and leave. The receiver can come and pick up letters later. They don't have to be there at the same time.
*/

/*
package main

import "fmt"

func main() {
	// Create a buffered channel with a capacity of 2
	message := make(chan string, 2)

	// Send operation do NOT block because the buffer is not full
	message <- "Hello"
	message <- "world"
	fmt.Println("Sent two messages without blocking")

	// Now, if we try to send a third, it will block (or deadlock
	// in this single-goroutine example), because the buffer is full.

	// message <- "extra" // This would cause a deadlock

	// Receive the messages
	fmt.Println(<-message)
	fmt.Println(<-message)

}
*/

/*

3. Directional Channels :

	- Directional channels are a compile-time feature in Go that restricts a channel to only one operation: either sending or receiving.

	- When you create a channel using make(chan int), you get a bidirectional channel, meaning you can both send to it (ch <- 10) and receive from it (<-ch).

*/

package main

import (
	"fmt"
	"sync"
	"time"
)

/*
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("Worker %d finished job %d\n", id, j)
		results <- j * 2
	}
}

func main() {
	// Unbuffered channel

	fmt.Println("\n Unbuffered channel")
	ch := make(chan string)
	go func() {
		ch <- "Hello from Gorutine" // send
	}()

	msg := <-ch //receive

	fmt.Println("Received ----> ", msg)

	// Buffered channel

	fmt.Println("\nBuffered channels start from here")
	bufch := make(chan int, 3)
	bufch <- 1
	bufch <- 2
	bufch <- 3

	fmt.Println("Buffer full")

	fmt.Println("Read-----", <-bufch)
	fmt.Println("Read-----", <-bufch)
	fmt.Println("Read-----", <-bufch)

	// Select statement
	fmt.Println("\n Select statment starts")
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Message from ch1"
	}()

	go func() {
		time.Sleep(3 * time.Second)
		ch2 <- "Message from ch2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-ch1:
			fmt.Println("ch1---- ", m1)

		case m2 := <-ch2:
			fmt.Println("ch2---- ", m2)
		}
	}

}
*/

//-------------------------------------------------------------------

func worker(url string, wg *sync.WaitGroup, resultChan chan string) {
	defer wg.Done()

	time.Sleep(time.Millisecond * 50)

	fmt.Printf("image processed: %s\n", url)

	resultChan <- url
}

func main() {
	var wg sync.WaitGroup
	resultChan := make(chan string, 5)

	startTime := time.Now()

	wg.Add(5)
	go worker("image_1.png", &wg, resultChan)
	go worker("image_2.png", &wg, resultChan)
	go worker("image_3.png", &wg, resultChan)
	go worker("image_4.png", &wg, resultChan)
	go worker("image_5.png", &wg, resultChan)

	wg.Wait()
	close(resultChan) // otherwise, ERROR: fatal error: all goroutines are asleep - deadlock!

	for result := range resultChan {
		fmt.Printf("received: %s\n", result)
	}

	fmt.Printf("received: %s\n", time.Since(startTime))

}
