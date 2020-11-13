package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile("h([a-z]+)king")
	fmt.Println(re.FindStringSubmatch("hacking hiking"))
	fmt.Println(re.FindStringSubmatch("licking"))
}
