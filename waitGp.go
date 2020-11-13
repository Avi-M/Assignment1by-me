package main

import (
	"fmt"
	"sync"
)

func f(wg *sync.WaitGroup) {
	fmt.Println("working.......")
	wg.Done()
}
func main() {

	var wg sync.WaitGroup
	wg.Add(1)
	go f(&wg)
	wg.Wait()
	fmt.Println("done Working.")

}
