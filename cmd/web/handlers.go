package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
	if r.Method != "GET" {
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	ts, err := template.ParseFiles("./ui/public/index.html")
	if err != nil {
		app.serverError(w, err)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) showContacts(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		app.clientError(w, http.StatusMethodNotAllowed)
	}
	//w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("show contacts"))
}

func (app *application) showSingleContact(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		app.clientError(w, http.StatusMethodNotAllowed)
	}
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	//use id
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Display specific contact with ID %d", id)
	w.Write([]byte("show contacts"))
}

func (app *application) addContact(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		app.clientError(w, http.StatusMethodNotAllowed)
	}
	w.Write([]byte("add contact"))
}
