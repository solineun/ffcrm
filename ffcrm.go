package main

import (
	"log"
	"net/http"
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
	w.Write([]byte("Snippet page\n"))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
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
