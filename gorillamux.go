package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Product struct {
	Id   int
	Name string
}

var products [10]Product

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Mux Rest !")
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	productsJson, _ := json.Marshal(products)
	fmt.Fprintf(w, string(productsJson))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/products", ProductsHandler).Methods("GET").Headers("Accept", "application/json")
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)

}
