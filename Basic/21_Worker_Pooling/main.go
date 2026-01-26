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

func worker(jobsChan chan string, wg *sync.WaitGroup, resultChan chan string) {
	defer wg.Done()

	for job := range jobsChan { // reading channel
		time.Sleep(time.Millisecond * 50)
		fmt.Printf("image processed : %s\n", job)
		resultChan <- job
	}

	fmt.Printf("Worker shutting down\n")
}

func main() {

	var wg sync.WaitGroup
	totalWorkers := 5

	resultChan := make(chan string, 50)
	jobsChan := make(chan string, len(jobs))

	startTime := time.Now()

	for i := 1; i < totalWorkers; i++ {
		wg.Add(1)
		go worker(jobsChan, &wg, resultChan)
	}

	go func() { // new goroutine for wait until all goroutines run
		wg.Wait()
		close(resultChan)
	}()

	// send the jobs
	for i := 0; i < len(jobs); i++ {
		jobsChan <- jobs[i]
	}

	close(jobsChan)

	for result := range resultChan {
		fmt.Printf("job completed : %s\n", result)
	}

	fmt.Printf("Total Time : %s\n", time.Since(startTime))

}
