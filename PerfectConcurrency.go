package main

import (
	"flag"
	"fmt"
	"runtime"
	"sync"
	"time"
)

var (
	n       = flag.Int("n", 1000, "Max number to search")
	workers = flag.Int("workers", runtime.GOMAXPROCS(0), "Number of workers")
	updates = flag.Duration("updates", 10*time.Second, "How often to print updates")
)

func isPerfect(n int) bool {
	var sum int
	for i := 1; i*i < n; i++ {
		if n%i != 0 {
			continue
		}
		sum += i
		if i != 1 {
			sum += n / i
		}
	}
	return sum == n
}

func main() {
	flag.Parse()

	progress := time.Tick(*updates)

	work := make(chan int, *workers*2)
	perfect := make(chan int)

	var wg sync.WaitGroup
	for i := 0; i < *workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for {
				candidate, ok := <-work
				if !ok {
					return
				}
				if isPerfect(candidate) {
					perfect <- candidate
				}
			}
		}()
	}

	go func() {
		for i := 2; i <= *n; {
			select {
			case work <- i:
				i++
			case <-progress:
				fmt.Printf("... %d\n", i)
			}
		}
		close(work)
		wg.Wait()
		close(perfect)
	}()

	start := time.Now()
	count := 0
	for {
		select {
		case n, ok := <-perfect:
			if !ok {
				goto done
			}
			fmt.Printf("Perfect: %d\n", n)
			count++
		}
	}

done:
	fmt.Printf("Found %d perfect numbers under %d in %s\n", count, *n, time.Since(start))
}
