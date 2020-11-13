package main

import (
	"errors"
	"fmt"
)

var ErrThreeNotFound = errors.New("error 3 is not found")

func main() {
	fmt.Println(DoSomething(1)) // succeeds! returns nil
	fmt.Println(DoSomething(2)) // returns a specific error message
	fmt.Println(DoSomething(3)) // returns an error variable
	fmt.Println(DoSomething(4)) // returns a simple error message
}

func DoSomething(someID int) error {
	switch someID {
	case 3:
		return ErrThreeNotFound
	case 2:
		return fmt.Errorf("this is an error with extra info: %d", someID)
	case 1:
		return nil
	}

	return errors.New("this is an error")
}
