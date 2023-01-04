package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	fileSever := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileSever))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView) // ?id=1
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Println("Starting server on port: 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
