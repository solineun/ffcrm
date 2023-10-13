package logic

import (
	"net/http"
)

type Logic interface {
	Logger
	Handler
}

type Logger interface {
	errLogger
	infoLogger
}

type errLogger interface {
	Fatal(err error)
	ServerError(http.ResponseWriter, error)
}

type infoLogger interface {
	Printf(format string, v any)
}

type Handler interface {
	customer	
}

type customer interface {
	Home(w http.ResponseWriter, r *http.Request)
}

type admin interface {
	AdminHome(w http.ResponseWriter, r *http.Request)
}