package main

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
  r := mux.NewRouter()

  r.HandleFunc("/books/{title}/page/{page}", HandleTitleAndPage).Methods("GET")

  http.ListenAndServe(":8080", r)
}

func HandleTitleAndPage(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)

  resp, _ := json.Marshal(vars)

  w.WriteHeader(200)
  w.Write(resp) 
}
