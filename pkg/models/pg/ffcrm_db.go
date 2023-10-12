package pg

import (
	"database/sql"
)

type DB interface {
	QueryRow(string, ...any) *sql.Row
	Query(string, ...any) (*sql.Rows, error)
	Select(interface{}, string, ...interface{}) error
}

type FFcrmDBadapt struct {
	db DB
}

func NewFFcrmDB(cfg Config) (*FFcrmDBadapt, error) {
	dbConn, err := newConn(cfg)
	if err != nil {
		return nil, err
	}
	
	return &FFcrmDBadapt{
		db: dbConn,
	}, nil
}
