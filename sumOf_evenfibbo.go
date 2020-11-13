package main

import (
	"fmt"
)

func fibonacci(x int) (int, string) {
	if x == 1 {
		return 1, "Yes"
	}
	if x == 2 {
		return 2, "Yes"
	}
	var sum1 int

	var sum2 int

	sum1, _ = fibonacci(x - 1)
	sum2, _ = fibonacci(x - 2)
	var sum = sum1 + sum2
	if sum > 4000000 {
		return sum, "No"
	}
	return sum, "Yes"
}

func main() {
	evensum := 0
	var sum int
	var ans string

	for i := 1; ; i++ {
		sum, ans = fibonacci(i)
		fmt.Println(i, sum)
		if ans == "No" {
			break
		} else {
			if sum%2 == 0 {

				evensum = evensum + sum
			}
		}

	}
	fmt.Println("Sum of even fibo is", evensum)
}
