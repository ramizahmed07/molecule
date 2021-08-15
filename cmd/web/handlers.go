package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	message := "Hello, Ramiz!"
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
