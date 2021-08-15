package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	message := "Hello, Ramiz!"

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte(message))
}

func showMolecule(w http.ResponseWriter, r *http.Request) {
	message := "Show molecule"
	w.Write([]byte(message))
}

func createMolecule(w http.ResponseWriter, r *http.Request) {
	message := "Creating a new molecule"
	w.Write([]byte(message))
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/molecule", showMolecule)
	mux.HandleFunc("/molecule/create", createMolecule)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
