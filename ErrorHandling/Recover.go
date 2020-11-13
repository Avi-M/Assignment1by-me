package main

import "fmt"

func foo() {
	panic("bar")
}

func bar() {
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Printf("Recovered with message %s\n", msg)
		}
	}()
	foo()
	fmt.Println("Never gets executed")
}

func main() {
	fmt.Println("Entering main")
	bar()
	fmt.Println("Exiting main the normal way")
}
