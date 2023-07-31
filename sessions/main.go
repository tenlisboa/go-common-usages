package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

var (
  key = []byte("super-secret-key")
  store = sessions.NewCookieStore(key)
)

func secret(w http.ResponseWriter, r *http.Request) {
  session, _ := store.Get(r, "auth-cookie")
  
  if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
    http.Error(w, "Forbidden", http.StatusForbidden)
    return
  }

  fmt.Fprintln(w, "Luke, I am your father!")
}

func login (w http.ResponseWriter, r * http.Request) {
  session, _ := store.Get(r, "auth-cookie")

  session.Values["authenticated"] = true

  session.Save(r, w)
}

func logout(w http.ResponseWriter, r *http.Request) {
  session, _ := store.Get(r, "auth-cookie")

  session.Values["authenticated"] = false

  session.Save(r, w)
}

func main() {
  http.HandleFunc("/secret", secret)
  http.HandleFunc("/login", login)
  http.HandleFunc("/logout", logout)

  http.ListenAndServe(":8080", nil)

}
