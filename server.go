package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	const port string = ":8080"
	router := mux.NewRouter()

	router.HandleFunc("/", func(response http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(response, "Server up and running")
	})
	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/posts", addPost).Methods("POST")

	log.Println("Server listening on port ", port)
	log.Fatalln(http.ListenAndServe(port, router))
}
