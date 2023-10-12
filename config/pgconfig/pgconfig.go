package pgconfig

import (
	"github.com/solineun/ffcrm/pkg/models/pg"
)

func GetPgConfig() pg.Config {
	return pg.Config {
		Addr: "localhost",
		Port:      5432,
		User: "postgres",
		Password: "pass",
		DB: "ffcrm",
	}
}