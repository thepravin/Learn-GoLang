/*
	- Worker Pooling is a pattern used to limit the number of goroutines running at the same time.

*/

package main

import (
	"fmt"
	"sync"
	"time"
)

var jobs = []string{
	"image_1.png",
	"image_2.png",
	"image_3.png",
	"image_4.png",
	"image_5.png",
	"image_6.png",
	"image_7.png",
	"image_8.png",
	"image_9.png",
	"image_10.png",
	"image_11.png",
	"image_12.png",
	"image_13.png",
	"image_14.png",
	"image_15.png",
	"image_16.png",
	"image_17.png",
	"image_18.png",
	"image_19.png",
	"image_20.png",
}

func worker(job string, wg *sync.WaitGroup, resultChan chan string) {
	defer wg.Done()

	time.Sleep(time.Millisecond * 50)

	fmt.Printf("processed : %s\n", job)

	resultChan <- job
}

func main() {

	var wg sync.WaitGroup

	resultChan := make(chan string, 50)

	startTime := time.Now()

	for _, job := range jobs {
		wg.Add(1)
		go worker(job, &wg, resultChan)
	}

	wg.Wait()
	close(resultChan)

	for result := range resultChan {
		fmt.Printf("received : %s\n", result)
	}

	fmt.Printf("Total Time : %s\n", time.Since(startTime))

}
