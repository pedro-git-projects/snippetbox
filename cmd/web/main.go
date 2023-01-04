package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	mux := http.NewServeMux()

	fileSever := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileSever))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView) // ?id=1
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Printf("Starting server on port%s", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
