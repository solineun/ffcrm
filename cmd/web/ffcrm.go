package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"solineun/ffcrm/pkg/models/pg"

	_ "github.com/lib/pq"
)

const url = "localhost:8080"

type application struct {
	errLog  *log.Logger
	infoLog *log.Logger
	orders *pg.OrderModel
}

func main() {
	var infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	var errLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	pgConfig := getConfig()
	db, err := openDb(pgConfig.format())
	if err != nil {
		errLog.Fatal(err)
	}
	defer db.Close()

	app := application{
		errLog:  errLog,
		infoLog: infoLog,
		orders: &pg.OrderModel{DB: db},
	}

	mux := app.routes()

	srv := &http.Server{
		Addr:     url,
		ErrorLog: errLog,
		Handler:  mux,
	}

	infoLog.Printf("starting web server on %s", url)
	err = srv.ListenAndServe()
	errLog.Fatal(err)
}

func openDb(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}