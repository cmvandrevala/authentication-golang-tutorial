package main

import (
  "net/http"
  "github.com/gorilla/mux"
)

func main() {
  r := mux.NewRouter()

  r.Handle("/", http.FileServer(http.Dir("./views/")))
  r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

  r.Handle("/status", notImplemented).Methods("GET")
  r.Handle("/products", notImplemented).Methods("GET")
  r.Handle("/products/{slug}/feedback", notImplemented).Methods("POST")

  http.ListenAndServe(":3000", r)
}

var notImplemented = http.HandlerFunc(
  func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Foo"))
  }
)
