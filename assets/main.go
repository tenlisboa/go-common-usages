package main

import (
	"net/http"
	"os"
	"path"
)

func main() {
  rootDir, _ := os.Getwd()
  currentDir := path.Join(rootDir, "assets")
  assetsDir := path.Join(currentDir, "static")

  fs := http.FileServer(http.Dir(assetsDir))

  http.Handle("/static/", http.StripPrefix("/static/", fs))

  http.ListenAndServe(":8080", nil)
}
