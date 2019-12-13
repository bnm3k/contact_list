package main

import (
	"database/sql"
	"flag"
	"fmt"
	"net/http"
	"runtime/debug"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func openDB(pgConnStr string) (*sql.DB, error) {
	//get db
	db, err := sql.Open("postgres", pgConnStr)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

type appVars struct {
	addr      string
	pgConnStr string
}

func getFlagArgs() appVars {
	//flag args
	hostArg := flag.String("host", "localhost", "Port to run http server on")
	portArg := flag.String("port", "4000", "Port to run http server on")
	dbPortArg := flag.String("dbport", "5432", "database port")
	dbHostArg := flag.String("dbhost", "localhost", "database host")
	dbUserArg := flag.String("dbuser", "", "database user")
	//dbPasswordArg := flag.String("dbpass", "", "database password")
	dbDBNameArg := flag.String("dbname", "contact_list", "database bnme")
	flag.Parse()
	pgConnStr := fmt.Sprintf("host=%s port=%s user=%s "+
		"dbname=%s sslmode=disable",
		*dbHostArg, *dbPortArg, *dbUserArg, *dbDBNameArg)
	addr := *hostArg + ":" + *portArg
	return appVars{
		addr:      addr,
		pgConnStr: pgConnStr,
	}
}
