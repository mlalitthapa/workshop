package main

import "fmt"

func main() {
	message := "hello"
	fmt.Println("message = ", message)
	valueType(message)
	fmt.Println("message = ", message)
	pointerType(&message)
	fmt.Println("message = ", message)
}

func valueType(msg string) {
	fmt.Println("value = ", msg)
	msg = "world"
}

func pointerType(msg *string) {
	fmt.Println("value = ", msg)
	*msg = "world"
}
