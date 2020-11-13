package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		fmt.Println("Running Anonymous fn")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Done.")
}
