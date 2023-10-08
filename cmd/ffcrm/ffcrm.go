package main

import (
	"github.com/solineun/ffcrm/internal/application"
	"github.com/solineun/ffcrm/internal/logic"
	"github.com/solineun/ffcrm/internal/server"
	"github.com/solineun/ffcrm/pkg/logadapt"
	"github.com/solineun/ffcrm/pkg/models/pg"
)

func main() {
	log := logadapt.NewLoggerAdapter()
	db, err := OpenDb(GetPgConfig().Format())
	if err != nil {
		log.Fatal(err)
		return
	}
	ffdb := pg.NewFFcrmDB(db)

	logic := logic.NewLogicAdapter(
		log,
		log,
		ffdb,
	)

	srv := server.NewServerAdapter(
		GetConfiguredSrv(),
	)

	app := application.NewApplication(
		logic, 
		srv,
	)
	app.Execute()
}