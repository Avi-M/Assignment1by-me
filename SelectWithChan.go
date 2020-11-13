package main

import (
	"fmt"
	"time"
)

func chatter(chatChannel chan<- string) {

	defer close(chatChannel)

	// loop five times and die
	for i := 1; i <= 5; i++ {
		time.Sleep(2 * time.Second) // sleep for 2 seconds
		chatChannel <- fmt.Sprintf("This is pass number %d of chatter", i)
	}
}

func main() {
	// Create the channel
	chatChannel := make(chan string, 1)

	go chatter(chatChannel)

	// This for loop keeps things going while the chatter is sleeping
	for {

		select {

		case spam, ok := <-chatChannel:

			if ok {
				fmt.Println(spam)
			} else {
				fmt.Println("Channel closed, exiting!")
				return
			}
		default:
			// print a line, then sleep for 1 second.
			fmt.Println("Nothing happened this second.")
			time.Sleep(1 * time.Second)
		}
	}
}
