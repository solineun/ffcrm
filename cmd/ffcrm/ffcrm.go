package main

import (
	"log"
	"net/http"
	"os"
	"github.com/solineun/ffcrm/pkg/models/pg"

	_ "github.com/lib/pq"
)

type application struct {
	url string
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

	app := application {
		url: "localhost:8080",
		errLog:  errLog,
		infoLog: infoLog,
		orders: pg.NewOrderModel(db),
	}

	mux := app.routes()

	srv := &http.Server {
		Addr:     app.url,
		ErrorLog: errLog,
		Handler:  mux,
	}

	infoLog.Printf("starting web server on %s", app.url)
	err = srv.ListenAndServe()
	errLog.Fatal(err)
}