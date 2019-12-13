package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/nagamocha3000/contact_list/pkg/models/postgres"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	contacts *postgres.ContactModel
}

func main() {
	appVars := getFlagArgs()

	//loggers
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	//db
	db, err := openDB(appVars.pgConnStr)
	if err != nil {
		errorLog.Fatal(err.Error())
	}
	defer db.Close()

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		contacts: &postgres.ContactModel{DB: db},
	}

	server := &http.Server{
		Addr:     appVars.addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("starting server on %s\n", appVars.addr)
	err = server.ListenAndServe()
	errorLog.Fatal(err)
}
