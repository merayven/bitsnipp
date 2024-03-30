package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello BitSnipp!"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet ID %d...", id)
}

func getSnippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Snippet Create"))
}

func postSnippetCreate(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)

	w.Write([]byte("Save a Snippet..."))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}/{$}", snippetView)
	mux.HandleFunc("GET /snippet/create", getSnippetCreate)
	mux.HandleFunc("POST /snippet/create", postSnippetCreate)

	log.Print("Starting server on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
