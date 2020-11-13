package main

import (
	"errors"
	"fmt"
)

type numbers struct {
	n1, n2, n3 int
}

func (n numbers) add() (int, error) {
	if n.n1 < 0 || n.n2 < 0 || n.n3 < 0 {
		return 0, errors.New("One of the numbers is negative")
	}
	return n.n1 + n.n2 + n.n3, nil
}

func main() {
	n := numbers{-11, 2, 3}
	if o, err := n.add(); err != nil {
		fmt.Println(err)
		// panic(err)
	} else {
		fmt.Println(o)
	}

	n1 := numbers{1, 2, 3}
	o, err := n1.add()
	fmt.Println(o)
	fmt.Println("Error is ", err)

}
