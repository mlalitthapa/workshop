package main

import "fmt"

func main() {
	ifCondition(55)
	switchCondtion("foo")
}

func ifCondition(num int) {
	if num%2 == 0 {
		fmt.Printf("%d is even \n", num)
	} else {
		fmt.Printf("%d is odd \n", num)
	}
}

func switchCondtion(pair string) {
	switch pair {
	case "foo":
		fmt.Printf("%s -> bar", pair)
	case "john":
		fmt.Printf("%s -> doe", pair)
	default:
		fmt.Printf("%s -> :(", pair)
	}
}
