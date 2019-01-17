package main

import (
	"html/template"
	"net/http"
)

type ToDo struct {
	Title string
	Done  bool
}

type PageData struct {
	Title string
	ToDos []ToDo
}

func main() {
	http.HandleFunc("/", handleIndex)

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/asset/", http.StripPrefix("/asset/", fs))

	http.ListenAndServe(":8000", nil)
}
func handleIndex(writer http.ResponseWriter, request *http.Request) {
	data := PageData{
		Title: "Task List",
		ToDos: []ToDo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: false},
		},
	}
	writer.Header().Set("Content-Type", "text/html")
	page := template.Must(template.ParseFiles("templates/index.html"))
	page.Execute(writer, data)
}
