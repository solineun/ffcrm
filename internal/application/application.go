package application

import (
	"github.com/solineun/ffcrm/internal/application/logic"
	"github.com/solineun/ffcrm/internal/application/server"
)

type Application struct {
	logic logic.Logic
	srv server.Server
}

func NewApplication(l logic.Logic, s server.Server) Application {
	return Application{
		logic: l,
		srv: s,
	}
}

func (app Application) Execute() {
	app.logic.Printf("starting web server%s", "")
	app.srv.HandleFunc("GET", "/", app.logic.Home)

	err := app.srv.ListenAndServe()
	if err != nil {
		app.logic.Fatal(err)
	}
}