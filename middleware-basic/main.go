package main

import (
	"fmt"
	"log"
	"net/http"
)

func logger(f http.HandlerFunc) http.HandlerFunc {
  return ensureJson(func(w http.ResponseWriter, r *http.Request) {
    log.Println(r.URL.Path)
    f(w, r)
  })
}

func ensureJson(f http.HandlerFunc) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Content-Type", "application/json")
    f(w, r)
  }
}

func foo(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "{\"message\": \"foo\"}")
}

func bar(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "{\"message\": \"bar\"}")
}

func main() {
  http.HandleFunc("/foo", logger(foo))
  http.HandleFunc("/bar", logger(bar))

  http.ListenAndServe(":8080", nil)
}
