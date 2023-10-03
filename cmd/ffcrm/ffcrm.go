package main

import (
	"net/http"
	"github.com/solineun/ffcrm/internal/applogic"
	lg "github.com/solineun/ffcrm/pkg/loggerimpl"
	"github.com/solineun/ffcrm/pkg/models/pg"
	"github.com/solineun/ffcrm/pkg/models/tmplcache"
)

var app *applogic.Application
const URL string = "localhost:8080"
var logger = lg.NewLogger()
var srv *http.Server

func main() {
	pgConfig := GetConfig()
	db, err := OpenDb(pgConfig.Format())
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	cache, err := tmplcache.NewTemplateCache("./ui/html/")
	if err != nil {
		logger.Fatal(err)
	}

	ffDb := pg.NewFfCrmDb(db)

	app = applogic.NewApplication(
		cache,
		logger,
		ffDb,
	)	

	mux := app.Routes()

	srv = &http.Server{
		Addr: URL,
		ErrorLog: logger.ErrLog,
		Handler: mux,
	}

	logger.Printf("starting web server on %s", URL)
	err = srv.ListenAndServe()
	logger.Fatal(err)
}

