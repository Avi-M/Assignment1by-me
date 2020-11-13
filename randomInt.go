package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var v [5]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ { //Generating  an array of random integers
		v[i] = rand.Intn(100)
	}
	fmt.Println(v)

	val := rand.Perm(9) //The perm function inside the rand package returns pseudo-random permutation of the integers 0 to n.
	fmt.Println(val)

}
