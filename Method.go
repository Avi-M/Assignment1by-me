package main

import "fmt"

type User struct {
	name string
}

func (u User) GetName() string {
	return u.name
}

func (u *User) SetName(newName string) {
	u.name = newName
}

func main() {
	var me User

	me.SetName("Avinash Maurya")
	fmt.Println("My name is", me.GetName())
}
