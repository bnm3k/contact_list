package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.Header().Set("Allow", "GET")
		http.NotFound(w, r)
		return
	}
	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")
		http.Error(w, "Method Not Allowed", 405)
	}
	ts, err := template.ParseFiles("./ui/public/index.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func showContacts(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")
		http.Error(w, "Method Not Allowed", 405)
	}
	//w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("show contacts"))
}

func showSingleContact(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")
		http.Error(w, "Method Not Allowed", 405)
	}
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	//use id
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Display specific contact with ID %d", id)
	w.Write([]byte("show contacts"))
}

func addContact(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method Not Allowed", 405)
	}
	w.Write([]byte("add contact"))
}
