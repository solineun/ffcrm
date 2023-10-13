package serveradapt

import (
	"fmt"
	"net"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ServerAdapter struct {
	handler *httprouter.Router
	srv *http.Server
}

// GetAddr implements Server.
func (sa *ServerAdapter) GetAddr() string {
	return sa.srv.Addr
}

// HandleFunc implements Server.
func (sa *ServerAdapter) HandleFunc(method string, pattern string, handler func(http.ResponseWriter, *http.Request)) {
	sa.handler.HandlerFunc(method, pattern, handler)
}

// ListenAndServe implements Server.
func (sa *ServerAdapter) ListenAndServe() error {
	sa.srv.Handler = sa.handler
	listener, err := net.Listen("tcp", ": 8080")
	if err != nil {
		return fmt.Errorf("listener init error: %v", err)
	}

	err = sa.srv.Serve(listener)
	return err
}

func NewServerAdapter(s *http.Server) *ServerAdapter {
	return &ServerAdapter{
		handler: httprouter.New(),
		srv: s,
	}
}
