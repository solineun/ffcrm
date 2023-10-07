package server

import (
	"net/http"
)

type Server interface {
	ListenAndServe() error
	HandleFunc(string, func(http.ResponseWriter, *http.Request))
	GetAddr() string
}

type ServerAdapter struct {
	srv *http.Server
}

// GetAddr implements Server.
func (sa *ServerAdapter) GetAddr() string {
	return sa.srv.Addr
}

// SetHandler implements Server.
func (sa *ServerAdapter) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(pattern, handler)
}

// ListenAndServe implements Server.
func (sa *ServerAdapter) ListenAndServe() error {
	err := sa.srv.ListenAndServe()
	return err
}

func NewServerAdapter(s *http.Server) *ServerAdapter {
	return &ServerAdapter{
		srv: s,
	}
}
