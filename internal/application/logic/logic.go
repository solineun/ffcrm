package logic

import (
	"net/http"
)

type Logic interface {
	Logger
	Handler
}

type Logger interface {
	ErrLogger
	InfoLogger
}

type ErrLogger interface {
	Fatal(err error)
	ServerError(http.ResponseWriter, error)
}

type InfoLogger interface {
	Printf(format string, v any)
}

type Handler interface {
	Home(w http.ResponseWriter, r *http.Request)
}