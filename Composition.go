package main

import (
	"fmt"
)

type Request struct {
	Resource string
}

type AuthenticatedRequest struct {
	Request
	Username, Password string
}

func main() {
	ar := new(AuthenticatedRequest)
	ar.Resource = "example.com/request"
	ar.Username = "bob"
	ar.Password = "P@ssw0rd"
	fmt.Printf("%v", ar)
}
