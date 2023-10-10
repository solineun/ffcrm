package main

import (
	"log"

	pgconf "github.com/solineun/ffcrm/config/pgconfig"
	srvconf "github.com/solineun/ffcrm/config/srvconfig"
	"github.com/solineun/ffcrm/internal/application"
	"github.com/solineun/ffcrm/internal/logicadapt"
	"github.com/solineun/ffcrm/internal/serveradapt"
	"github.com/solineun/ffcrm/pkg/models/pg"
)

func main() {	
	db, err := pgconf.OpenDb(pgconf.GetPgConfig().Format())
	if err != nil {
		log.Fatal(err)
		return
	}
	ffdb := pg.NewFFcrmDB(db)
	
	logic := logicadapt.NewLogicAdapter(ffdb)
	srv := serveradapt.NewServerAdapter(srvconf.GetConfiguredSrv())

	app := application.NewApplication(
		logic, 
		srv,
	)
	app.Execute()
}