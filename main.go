package main

import (
  "net/http"
  "github.com/gorilla/mux"
)

type Product struct {
  id int
  name string
  slug string
  description string
}

var productList = []Product {
  Product{id: 0, name: "hover board", slug: "slug-one", description: "Some description"},
  Product{id: 1, name: "Hover Shooters", slug: "hover-shooters", description : "Shoot your way to the top on 14 different hoverboards"},
  Product{id: 2, name: "Ocean Explorer", slug: "ocean-explorer", description : "Explore the depths of the sea in this one of a kind underwater experience"},
  Product{id: 3, name: "Dinosaur Park", slug : "dinosaur-park", description : "Go back 65 million years in the past and ride a T-Rex"},
  Product{id: 4, name: "Cars VR", slug : "cars-vr", description: "Get behind the wheel of the fastest cars in the world."},
  Product{id: 5, name: "Robin Hood", slug: "robin-hood", description : "Pick up the bow and arrow and master the art of archery"},
  Product{id: 6, name: "Real World VR", slug: "real-world-vr", description : "Explore the seven wonders of the world in VR"},
}

func main() {
  r := mux.NewRouter()

  r.Handle("/", http.FileServer(http.Dir("./views/")))
  r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

  r.Handle("/status", StatusHandler).Methods("GET")
  r.Handle("/products", NotImplemented).Methods("GET")
  r.Handle("/products/{slug}/feedback", NotImplemented).Methods("POST")

  http.ListenAndServe(":3000", r)
}

var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
  w.Write([]byte("Not Implemented"))
})

var StatusHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
  w.Write([]byte("API is up and running"))
})
