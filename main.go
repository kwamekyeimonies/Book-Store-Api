package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kwamekyeimonies/Book-Store-Api/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}

//1:30 / 8:24 11 projects
