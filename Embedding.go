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
	ResourceFormatter
}

type ResourceFormatter struct{}

func (r *ResourceFormatter) FormatHTTP(resource string) string {
	return fmt.Sprintf("http://%s", resource)
}

func (r *ResourceFormatter) FormatHTTPS(resource string) string {
	return fmt.Sprintf("https://%s", resource)
}

func main() {
	ar := new(AuthenticatedRequest)
	ar.Resource = "www.example.com/request"
	ar.Username = "bob"
	ar.Password = "P@ssw0rd"

	println(ar.FormatHTTP(ar.Resource))
	println(ar.FormatHTTPS(ar.Resource))

	fmt.Printf("%#v", ar)
}
