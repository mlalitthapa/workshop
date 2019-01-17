package main

import (
	"encoding/json"
	"fmt"
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

type FormData struct {
	Name, Email, Message string
}

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/json", handleJsonResponse)
	http.HandleFunc("/form", showFormHandler)
	http.HandleFunc("/form/post", postFormHandler)

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/asset/", http.StripPrefix("/asset/", fs))

	http.ListenAndServe(":8000", nil)
}

func postFormHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Redirect(writer, request, "/form", 302)
		return
	}

	formData := FormData{
		Name: request.FormValue("name"),
		Email: request.FormValue("email"),
		Message: request.FormValue("message"),
	}

	json.NewEncoder(writer).Encode(formData)
}

func showFormHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html")
	form := template.Must(template.ParseFiles("templates/form.html"))
	form.Execute(writer, nil)
}

func handleJsonResponse(writer http.ResponseWriter, request *http.Request) {
	err := json.NewEncoder(writer).Encode(getPageData())
	if err != nil {
		fmt.Println(err)
	}
}

func handleIndex(writer http.ResponseWriter, request *http.Request) {
	data := getPageData()
	writer.Header().Set("Content-Type", "text/html")
	page := template.Must(template.ParseFiles("templates/index.html"))
	page.Execute(writer, data)
}

func getPageData() PageData {
	return PageData{
		Title: "Task List",
		ToDos: []ToDo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: false},
		},
	}
}
