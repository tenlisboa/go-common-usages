package main

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

type ContactDetails struct {
  Email   string
  Subject string
  Message string
}

func main() {
  rootDir, _ := os.Getwd()
  currDir := filepath.Join(rootDir, "forms")
  formsFilepath := filepath.Join(currDir, "forms.html")
  tmpl := template.Must(template.ParseFiles(formsFilepath))

  http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
      tmpl.Execute(w, nil)
      return
    }

    details := ContactDetails{
      Email:   r.FormValue("email"),
      Subject: r.FormValue("subject"),
      Message: r.FormValue("message"),
    }

    _ = details

    tmpl.Execute(w, struct{Success bool}{true})
  })

  http.ListenAndServe(":8080", nil)
}
