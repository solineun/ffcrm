package main

import (
	"log"
	"net/http"
	"os"
)

const url = "localhost:8080"

type application struct {
	errLog  *log.Logger
	infoLog *log.Logger
}

func main() {
	var infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	var errLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := application{
		errLog:  errLog,
		infoLog: infoLog,
	}

	mux := app.routes()

	srv := &http.Server{
		Addr:     url,
		ErrorLog: errLog,
		Handler:  mux,
	}

	infoLog.Printf("starting web server on %s", url)
	err := srv.ListenAndServe()
	errLog.Fatal(err)
}