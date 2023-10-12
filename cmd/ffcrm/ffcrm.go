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
	dbConf := pgconf.GetPgConfig()
	ffdb, err := pg.NewFFcrmDB(dbConf)
	if err != nil {
		log.Fatal(err)
	}
		
	logic := logicadapt.NewLogicAdapter(ffdb)
	srv := serveradapt.NewServerAdapter(srvconf.GetConfiguredSrv())

	app := application.NewApplication(
		logic, 
		srv,
	)
	app.Execute()
}