package main

import (
	"log"
	"net/http"
	"time"
)

type Middleware func (http.HandlerFunc) http.HandlerFunc

func Logger() Middleware {
  return func(f http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
      start := time.Now()
      defer func() {log.Println(r.URL.Path, time.Since(start))}()

      f(w, r)
    }
  }
}

func Method(m string) Middleware {
  return func(f http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
      if r.Method != m {
        http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
        return
      }

      f(w, r)
    }
  }
}

func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
  for _, m := range middlewares {
    f = m(f)
  }

  return f
}

func Hello(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Hello World"))
}

func main() {
  http.HandleFunc("/", Chain(Hello, Method("GET"), Logger()))

  http.ListenAndServe(":8080", nil)
}

