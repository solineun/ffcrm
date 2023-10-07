package pg

import "database/sql"

type FFcrmDBadapt struct {
	orderModel
}

var ffcrmDB *sql.DB

func NewFFcrmDB(db *sql.DB) *FFcrmDBadapt {
	ffcrmDB = db	
	return &FFcrmDBadapt{}
}