package main

import (
	"fmt"
)

// Define AuthorizationError as composite literal
type AuthorizationError string

// Implement the error interface
// In this case, I simply return the underlying string
func (e AuthorizationError) Error() string {
	return string(e)
}

func main() {
	fmt.Println(DoSomething(1)) // succeeds! returns nil
	fmt.Println(DoSomething(2)) // returns an error message
}

func DoSomething(someID int) error {
	if someID != 1 {
		return AuthorizationError("Action not allowed!")
	}

	return nil
}
