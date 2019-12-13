package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	staticFilesHandle := http.StripPrefix("/static/", http.FileServer(http.Dir("./ui/public/")))
	mux.HandleFunc("/", home)
	mux.Handle("/static/", staticFilesHandle)
	mux.HandleFunc("/contact", showSingleContact)
	mux.HandleFunc("/contacts", showContacts)
	mux.HandleFunc("/contact/add", addContact)

	log.Println("starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
