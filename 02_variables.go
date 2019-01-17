package main

import "fmt"

var globalVariable string

func main() {

	globalVariable = "This is global variable"

	var a string
	var b = "string"

	fmt.Printf("a = %s \nb = %s \n", a, b)

	x := 1
	y := 0.001
	fmt.Printf("x = %d\ny = %.3f \n", x, y)

	var t bool
	var f = false
	fmt.Printf("t = %t \nf = %t \n", t, f)

	otherFunction()
}

func otherFunction() {
	fmt.Printf("global variable = %s", globalVariable)
}