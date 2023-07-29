package main

import (
	"html/template"
	"net/http"
	"os"
	"path"
)

type Todo struct {
  Title string
  Done  bool
}

type TodoPageData struct {
  PageTitle string
  Todos     []Todo
}

func main() {
  dir, _ := os.Getwd()
  templateDir := path.Join(dir, "templates")

  data := TodoPageData{
    PageTitle: "My TODO list",
    Todos: []Todo{
      {Title: "Task 1", Done: false},
      {Title: "Task 2", Done: true},
      {Title: "Task 3", Done: true},
    },
  }

  layoutTemplate := path.Join(templateDir, "layout.html") 
  tmpl := template.Must(template.ParseFiles(layoutTemplate))

  http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
    tmpl.Execute(w, data)
  })

  http.ListenAndServe("localhost:8080", nil)
}
