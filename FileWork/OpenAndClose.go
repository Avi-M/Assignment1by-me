package main

import (
	"log"
	"os"
)

func main() {

	file, err := os.Open("test1.txt")
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	file, err = os.OpenFile("test1.txt", os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	// os.O_RDONLY // Read only
	// os.O_WRONLY // Write only
	// os.O_RDWR // Read and write
	// os.O_APPEND // Append to end of file
	// os.O_CREATE // Create is none exist
	// os.O_TRUNC // Truncate file when opening
}
