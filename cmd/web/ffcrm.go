package main

import (
	"log"
	"net/http"
)

const url = "localhost:8080"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Printf("starting web server on %s", url)
	err := http.ListenAndServe(url, mux)
	log.Fatal(err)
}
