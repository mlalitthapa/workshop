package main

import "fmt"

func main() {
	sum := sum(1, 2, 3, 4, 5)
	fmt.Println(sum)

	nums := []int{2, 5, 7, 9, 15}
	average := average(nums, total)
	fmt.Println(average)
}

// Variadic function
func sum(args ...int) int {
	sum := 0
	for _, n := range args{
		sum += n
	}
	return sum
}

func total(nums []int) int {
	var sum int
	for _, n := range nums {
		sum += n
	}
	return sum
}

func average(nums []int, cb func([]int) int) float64 {
	total := cb(nums)
	return float64(total) / float64(len(nums))
}
