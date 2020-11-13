package main

import (
	"fmt"
	"time"
)

func chatter(chatChannel chan<- string) {

	time.Sleep(5 * time.Second) // sleep for 5 seconds
	chatChannel <- fmt.Sprintf("This is pass number %d of chatter", 1)
}

func main() {

	chatChannel := make(chan string)

	defer close(chatChannel)

	go chatter(chatChannel)

	select {

	case spam := <-chatChannel:
		fmt.Println(spam)

	case <-time.After(3 * time.Second):
		fmt.Println("Ain't no time for that!")
	}
}
