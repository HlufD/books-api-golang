package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/HlufD/books_api/routes"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Bookd api")
	r := mux.NewRouter()
	routes.BookRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":4000", r))
}
