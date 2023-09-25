package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const url = "localhost:8080"

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Homepage\n"))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Snippet page with id=%d \n", id)
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "405 Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("New snippet form\n"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Printf("starting web server on %s", url)
	err := http.ListenAndServe(url, mux)
	log.Fatal(err)
}