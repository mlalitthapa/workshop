package main

import "fmt"

func main() {
	arrays()
}

func arrays() {
	arr := [5]int{1, 2, 3, 4}
	arr[4] = 5

	for _, a := range arr {
		fmt.Println(a)
	}
}
