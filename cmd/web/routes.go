package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	getStaticFiles := http.StripPrefix("/static/", http.FileServer(http.Dir("./ui/public/")))
	mux.Handle("/static/", getStaticFiles)
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/contact", app.showSingleContact)
	mux.HandleFunc("/contacts", app.showContacts)
	mux.HandleFunc("/contact/add", app.addContact)
	return mux
}
