package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/solineun/ffcrm/internal/applogic"
	lg "github.com/solineun/ffcrm/pkg/loggerimpl"
	"github.com/solineun/ffcrm/pkg/models/pg"
	"github.com/solineun/ffcrm/pkg/models/tmplcache"

	_ "github.com/lib/pq"
)

var app *applogic.Application
const URL string = "localhost:8080"
var logger = lg.NewLogger(nil, nil)
var srv *http.Server

func main() {
	pgConfig := getConfig()
	db, err := openDb(pgConfig.format())
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
		ErrorLog: lg.DefaultErr,
		Handler: mux,
	}

	logger.Printf("starting web server on %s", URL)
	err = srv.ListenAndServe()
	logger.Fatal(err)
}

type pgConfig struct {
	host string
	port uint16
	user string
	passwd string
	dbName string
}

func getConfig() pgConfig {
	return pgConfig {
		host: "localhost",
		port:      5432,
		user: "postgres",
		passwd: "pass",
		dbName: "ffcrm",
	}
}

func (pgc pgConfig) format() string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    pgc.host, pgc.port, pgc.user, pgc.passwd, pgc.dbName)	
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