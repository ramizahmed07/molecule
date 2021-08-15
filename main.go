package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
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
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	message := fmt.Sprintf("Show molecule with ID %d", id)
	w.Write([]byte(message))
}

func createMolecule(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		const code = 405
		http.Error(w, "Method Not Allowed", code)
		return
	}
	message := "Creating a new molecule\n"
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
