package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleIndex)

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/asset/", http.StripPrefix("/asset/", fs))

	http.ListenAndServe(":8000", nil)
}
func handleIndex(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello World")
}
