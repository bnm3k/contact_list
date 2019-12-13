package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	hostArg := flag.String("host", "localhost", "Port to run http server on")
	portArg := flag.String("port", "4000", "Port to run http server on")
	flag.Parse()
	addr := *hostArg + ":" + *portArg

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	server := &http.Server{
		Addr:     addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("starting server on %s\n", addr)
	err := server.ListenAndServe()
	errorLog.Fatal(err)
}
