package pg

import "database/sql"

type FfcrmDbImpl struct {
	orderModel
}

var ffcrmDb *sql.DB

func NewFfCrmDb(db *sql.DB) *FfcrmDbImpl {
	ffcrmDb = db	
	return &FfcrmDbImpl{}
}