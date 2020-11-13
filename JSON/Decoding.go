package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	encodedValue := []byte(`{"London":18,"Rome":30}`)

	// generic storage for the decoded JSON
	var data map[string]int

	// decode the value into data

	err := json.Unmarshal(encodedValue, &data)

	// check if the decoding is successful
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
