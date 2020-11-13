package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	numPiecesOfWork := 20
	numWorkers := 5

	workCh := make(chan int)
	wg := &sync.WaitGroup{}

	// Start workers
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go worker(workCh, wg)
	}

	// Send work
	for i := 0; i < numPiecesOfWork; i++ {
		work := i % 10 // invent some work
		workCh <- work
	}

	// Tell workers that no more work is coming
	close(workCh)

	// Wait for workers to finish
	wg.Wait()

	fmt.Println("done")
}

func worker(workCh <-chan int, wg *sync.WaitGroup) {
	defer wg.Done() // will call wg.Done() right before returning

	for work := range workCh { // will wait for work until workCh is closed
		doWork(work)
	}
}

func doWork(work int) {
	time.Sleep(time.Duration(work) * time.Millisecond)
	fmt.Println("slept for", work, "milliseconds")
}
