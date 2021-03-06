// pingpong
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func hit(ch chan int) {
	i := <-ch
	if i%7 == 0 {
		fmt.Println("Ping lost")
		return
	}
	fmt.Printf("Ping Received %d \n", i)
	i = i*3 + 29
	fmt.Printf("Ping Hit %d \n", i)

	go receive(ch)
	ch <- i
	time.Sleep(100 * time.Millisecond)
}

func receive(ch chan int) {
	i := <-ch
	if i%9 == 0 {
		fmt.Println("Pong lost")
		return
	}
	fmt.Printf("Pong Received %d \n", i)
	go hit(ch)
	time.Sleep(100 * time.Millisecond)
	i = i + 1
	ch <- i
	fmt.Printf("Pong Hit %d \n", i)

}

func main() {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(100)
	var ch = make(chan int)
	fmt.Printf("Serving with %d \n", r)
	go hit(ch)
	ch <- r
	time.Sleep(2 * time.Minute)

}
