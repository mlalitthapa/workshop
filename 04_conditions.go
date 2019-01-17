package main

import "fmt"

func main() {
	ifCondition(55)
}

func ifCondition(num int) {
	if num % 2 == 0 {
		fmt.Printf("%d is even \n", num)
	} else {
		fmt.Printf("%d is odd \n", num)
	}
}
