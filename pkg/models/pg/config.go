package pg

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

)

type Config struct {
	Addr string
	Port uint16
	User string
	Password string
	DB string
}

func newConn(cfg Config) (*sqlx.DB, error) {
	dataSource := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable", 
		cfg.User, cfg.Password, cfg.Addr, cfg.Port, cfg.DB)
	
	conn, err := sqlx.Connect("postgres", dataSource)
	if err != nil {
		return nil, fmt.Errorf("sqlx connect: %v", err)
	}

	err = conn.Ping()
	if err != nil {
		return nil, fmt.Errorf("ping failed: %v", err)
	}

	return conn, nil
}