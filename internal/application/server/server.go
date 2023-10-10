package server

import (
	"net/http"
)

type Server interface {
	ListenAndServe() error
	HandleFunc(string, func(http.ResponseWriter, *http.Request))
	GetAddr() string
}
