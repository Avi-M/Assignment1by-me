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
	// data to encode
	city := City{Name: "Rome", Temperature: 30}

	// encode the data
	bytes, err := json.Marshal(city)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))

}
