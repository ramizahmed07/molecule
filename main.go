package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	message := "Hello, Ramiz!"
	w.Write([]byte(message))
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
