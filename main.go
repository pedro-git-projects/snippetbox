package main

import (
	"log"
	"net/http"
)

// Home handler
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, snippetbox"))
}

func main() {
	mux := http.NewServeMux()
	// registers the function home as the handler for the pattern /
	mux.HandleFunc("/", home)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux) // starts a new Web Server
	log.Fatal(err)
}
