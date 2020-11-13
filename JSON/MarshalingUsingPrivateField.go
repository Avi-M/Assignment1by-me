package main

import (
	"encoding/json"
	"fmt"
)

type MyStruct struct {
	uuid string
	Name string
}

func main() {
	myStruct := MyStruct{uuid: "PROPER-UUID-STRING", Name: "Some Proper Name"}
	j, err := json.Marshal(myStruct)
	if err != nil {
		fmt.Println("Something went wrong")
		return
	}
	fmt.Println(string(j))
}
