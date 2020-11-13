package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	someValue := []string{"foo", "bar"}
	data, err := json.Marshal(someValue)

	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

}
