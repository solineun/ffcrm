package logger

import "net/http"

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

