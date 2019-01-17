package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleIndex)
	http.ListenAndServe(":8000", nil)
}
func handleIndex(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello World")
}
