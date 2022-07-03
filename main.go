package main

import (
	"log"
	"net/http"
)

/* Handlers */
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, snippetbox"))

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("showing snippet"))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")          // will indicate that POST is the allowed method
		http.Error(w, "Method Not Allowed", 405) // calls w.WriteHeader() and w.Write() under the hood
		return
	}
	w.Write([]byte("creating snippet"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
