package main

import "fmt"

func main() {
	total := sum(1, 2, 3, 4, 5)
	fmt.Println(total)
}

// Variadic function
func sum(args ...int) int {
	sum := 0
	for _, n := range args{
		sum += n
	}
	return sum
}
