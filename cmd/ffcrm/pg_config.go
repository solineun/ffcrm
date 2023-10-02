package main

import (
	"database/sql"
	"fmt"
)

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