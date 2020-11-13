package main

import (
	"encoding/json"
	"fmt"
)

type Company struct {
	Name     string `json:"name"`
	Location string `json:"location,omitempty"`
	Revenue  int    `json:"-"`
	sales    int
}

func main() {

	company := Company{Name: "Stackoverflow", Location: "New York"}
	bytes, err := json.Marshal(company)
	if err != nil {
		panic(err)
	}
	company.Revenue = 1000000
	company.sales = 1000000
	fmt.Println(string(bytes))

	company = Company{Name: "StackExchange"}
	bytes, err = json.Marshal(company)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))

}
