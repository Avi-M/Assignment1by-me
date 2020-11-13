// Golang program to illustrate the
// string matching using regexp
// in-built functions
package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {

	re, _ := regexp.Compile("Kloud")

	str := "I love KloudforKloudOne"

	match := re.FindStringIndex(str)
	fmt.Println(match)

	str2 := "I love computer science"

	match2 := re.FindStringIndex(str2)
	fmt.Println(match2)

	re2, _ := regexp.Compile("[0-9]+-v.*g")

	match3 := re2.FindString("20024-vani_gupta")
	fmt.Println(match3)

	match4 := re.FindAllStringSubmatchIndex("I'am a geek at"+
		" geeksforgeeks", -1)
	fmt.Println(match4)

	re3, _ := regexp.Compile(" ")
	match5 := re3.ReplaceAllString("All I do"+
		" is code everytime.", "+")
	fmt.Println(match5)

	re4, _ := regexp.Compile("[aeiou]+")
	match6 := re4.ReplaceAllStringFunc("All I do"+
		" is code everytime.", strings.ToUpper)
	fmt.Println(match6)
}
