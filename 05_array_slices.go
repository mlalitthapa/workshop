package main

import "fmt"

func main() {
	arrays()
	slices()
}

func arrays() {
	arr := [5]int{1, 2, 3, 4}
	arr[4] = 5

	for _, a := range arr {
		fmt.Println(a)
	}
}

func slices() {
	var slice []string
	slice = append(slice, "first item", "second item")
	fmt.Println(slice[0], slice[1])
}
