package main

import (
	"encoding/json"
	"fmt"
)

type City struct {
	Name        string `json:"name"`
	Temperature int    `json:"temp"`
}

func main() {
	// data to decode
	bytes := []byte(`{"name":"Rome","temp":30}`)

	// initialize the container for the decoded data
	var city City

	// decode the data
	// notice the use of &city to pass the pointer to city
	err := json.Unmarshal(bytes, &city)

	// check if the decoding is successful
	if err != nil {
		panic(err)
	}

	fmt.Println(city)
}
