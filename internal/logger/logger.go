package logger

import "net/http"

type ErrLogger interface {
	ServerError(w http.ResponseWriter, err error)
	ClientError(w http.ResponseWriter, status int)
}

type InfoLogger interface {
	Printf(format string, v any)
}