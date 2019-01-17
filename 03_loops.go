package main

import "fmt"

func main() {
	i := 1

	for i <= 3 {
		fmt.Println(i)
		i++
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	numbers := []int{1, 2, 3, 4, 5}
	for i, n := range numbers {
		fmt.Println(i, n)
	}
}
